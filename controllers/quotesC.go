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
