package controllers

import (
	"context"
	"net/http"
	"promg/configs"
	"promg/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var quoteCollection *mongo.Collection = configs.GetCollection(configs.DB, "quotes")

// *********************************************************      items
func NewQuote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newQuote models.Quote
		defer cancel()

		//validate the request body
		if err := c.ShouldBindJSON(&newQuote); err != nil {
			println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}
		folio := checkQuoteFolio(newQuote.CID)
		newQuote.Folio = folio + 1
		_objId, _ := primitive.ObjectIDFromHex(newQuote.CID)
		updateFolio(newQuote.Folio, _objId)

		_, err := quoteCollection.InsertOne(ctx, newQuote)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		//return Success At creation
		c.JSON(http.StatusCreated, gin.H{"data": newQuote.Folio})
	}
}

func GetQuotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		var quotes []models.Quote
		defer cancel()
		filter := bson.D{{Key: "cid", Value: cid}}
		opts := options.Find().SetSort(bson.D{{Key: "folio", Value: -1}})

		results, err := quoteCollection.Find(ctx, filter, opts)

		if err != nil {
			println(err.Error())
		}

		for results.Next(context.TODO()) {
			var elem models.Quote
			err := results.Decode(&elem)

			if err != nil {

			}
			quotes = append(quotes, elem)

		}

		if err = results.All(context.TODO(), &quotes); err != nil {
			println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		if len(quotes) != 0 {
			c.JSON(http.StatusOK, gin.H{"data": quotes})
		} else {

			c.JSON(http.StatusOK, gin.H{"data": false})
		}

	}
}

func UpdateQuote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		oid := c.GetHeader("oid")
		var updateQuote models.Quote
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&updateQuote); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}
		ObjectId, _ := primitive.ObjectIDFromHex(oid)
		quoteUpdated, err := quoteCollection.UpdateOne(ctx, bson.M{"cid": updateQuote.CID, "_id": ObjectId}, bson.M{"$set": updateQuote})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": quoteUpdated.MatchedCount})

	}
}

func DeleteQuote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		oid := c.GetHeader("oid")
		cid := c.GetHeader("cid")

		defer cancel()

		ObjectId, _ := primitive.ObjectIDFromHex(oid)
		deleteResult, err := quoteCollection.DeleteOne(ctx, bson.M{"cid": cid, "_id": ObjectId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		if deleteResult.DeletedCount != 0 {
			c.JSON(http.StatusOK, gin.H{"data": true})
		} else {
			c.JSON(http.StatusNoContent, gin.H{"data": false})
		}

	}
}

func CopyQuote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var copyQuote models.Quote
		defer cancel()

		//validate the request body
		if err := c.ShouldBindJSON(&copyQuote); err != nil {
			println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}
		folio := checkQuoteFolio(copyQuote.CID)
		copyQuote.Folio = folio + 1
		_objId, _ := primitive.ObjectIDFromHex(copyQuote.CID)
		updateFolio(copyQuote.Folio, _objId)

		_, err := quoteCollection.InsertOne(ctx, copyQuote)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		//return Success At creation
		c.JSON(http.StatusCreated, gin.H{"data": copyQuote.Folio})
	}
}

func checkQuoteFolio(cid string) int {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	objId, _ := primitive.ObjectIDFromHex(cid)
	var company models.Company
	var result models.CompanyDb
	defer cancel()

	err := companyCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&result)
	if err != nil {
		// TODO implementar LOGS
	}
	company = models.Company(result)
	return company.Conf.QFolio

}

func updateFolio(folio int, objId primitive.ObjectID) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	companyUpdated, err := companyCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": bson.M{"conf.qfolio": folio}})
	defer cancel()

	if err != nil {
		return false
	} else {
		if companyUpdated.MatchedCount == 1 {
			return true
		} else {
			return false
		}
	}
}
