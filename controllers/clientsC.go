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

var clientCollection *mongo.Collection = configs.GetCollection(configs.DB, "clients")

// *********************************************************      clients
func NewClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newClient models.Client
		defer cancel()

		//validate the request body
		if err := c.ShouldBindJSON(&newClient); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}
		_, err := clientCollection.InsertOne(ctx, newClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		//return Success At creation
		c.JSON(http.StatusCreated, gin.H{"data": "created"})
	}
}

func GetClients() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		var clients []models.Client
		defer cancel()
		filter := bson.D{{Key: "cid", Value: cid}}

		results, err := clientCollection.Find(ctx, filter)
		if err != nil {
			println(err.Error())
		}

		for results.Next(context.TODO()) {
			var elem models.Client
			err := results.Decode(&elem)
			if err != nil {
			}
			clients = append(clients, elem)
		}

		if err = results.All(context.TODO(), &clients); err != nil {
			println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": clients})
	}
}

func UpdateClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		oid := c.GetHeader("oid")
		var updateClient models.Client
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&updateClient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}
		ObjectId, _ := primitive.ObjectIDFromHex(oid)
		clientUpdated, err := clientCollection.UpdateOne(ctx, bson.M{"cid": updateClient.CID, "_id": ObjectId}, bson.M{"$set": updateClient})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": clientUpdated.MatchedCount})

	}
}

func DeleteClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		oid := c.GetHeader("oid")
		cid := c.GetHeader("cid")

		defer cancel()

		ObjectId, _ := primitive.ObjectIDFromHex(oid)
		deleteResult, err := clientCollection.DeleteOne(ctx, bson.M{"cid": cid, "_id": ObjectId})
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
