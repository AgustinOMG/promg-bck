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
)

var itemCollection *mongo.Collection = configs.GetCollection(configs.DB, "items")

// *********************************************************      items
func NewItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newItem models.Item
		defer cancel()
		println(newItem.Code)
		//validate the request body
		if err := c.ShouldBindJSON(&newItem); err != nil {
			println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}
		_, err := itemCollection.InsertOne(ctx, newItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		//return Success At creation
		c.JSON(http.StatusCreated, gin.H{"data": "created"})
	}
}

func GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		var items []models.Item
		defer cancel()
		filter := bson.D{{Key: "cid", Value: cid}}

		results, err := itemCollection.Find(ctx, filter)
		if err != nil {
			println(err.Error())
		}

		for results.Next(context.TODO()) {
			var elem models.Item
			err := results.Decode(&elem)
			if err != nil {
			}
			items = append(items, elem)
		}

		if err = results.All(context.TODO(), &items); err != nil {
			println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": items})
	}
}

func UpdateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		oid := c.GetHeader("oid")
		var updateItem models.Item
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&updateItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}
		ObjectId, _ := primitive.ObjectIDFromHex(oid)
		itemUpdated, err := itemCollection.UpdateOne(ctx, bson.M{"cid": updateItem.CID, "_id": ObjectId}, bson.M{"$set": updateItem})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": itemUpdated.MatchedCount})

	}
}

func DeleteItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		oid := c.GetHeader("oid")
		cid := c.GetHeader("cid")

		defer cancel()

		ObjectId, _ := primitive.ObjectIDFromHex(oid)
		deleteResult, err := itemCollection.DeleteOne(ctx, bson.M{"cid": cid, "_id": ObjectId})
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
