package controllers

import (
	"context"
	"net/http"
	"promg/configs"
	"promg/models"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

		_, err := quoteCollection.InsertOne(ctx, newQuote)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		//return Success At creation
		c.JSON(http.StatusCreated, gin.H{"data": "created"})
	}
}

func GetQuotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		var quotes []models.Quote
		defer cancel()
		filter := bson.D{{Key: "cid", Value: cid}}

		results, err := quoteCollection.Find(ctx, filter)
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

		c.JSON(http.StatusOK, gin.H{"data": quotes})
	}
}
