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

var supplierCollection *mongo.Collection = configs.GetCollection(configs.DB, "suppliers")

// *********************************************************      Suppliers
func NewSupplier() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newSupplier models.Supplier
		defer cancel()

		//validate the request body
		if err := c.ShouldBindJSON(&newSupplier); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}
		_, err := supplierCollection.InsertOne(ctx, newSupplier)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		//return Success At creation
		c.JSON(http.StatusCreated, gin.H{"data": "created"})
	}
}

func GetSuppliers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		var suppliers []models.Supplier
		defer cancel()
		filter := bson.D{{Key: "cid", Value: cid}}

		results, err := supplierCollection.Find(ctx, filter)
		if err != nil {
			println(err.Error())
		}

		for results.Next(context.TODO()) {
			var elem models.Supplier
			err := results.Decode(&elem)
			if err != nil {
			}
			suppliers = append(suppliers, elem)
		}

		if err = results.All(context.TODO(), &suppliers); err != nil {
			println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": suppliers})
	}
}

func UpdateSupplier() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		oid := c.GetHeader("oid")
		var updateSupplier models.Supplier
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&updateSupplier); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}
		ObjectId, _ := primitive.ObjectIDFromHex(oid)
		SupplierUpdated, err := supplierCollection.UpdateOne(ctx, bson.M{"cid": updateSupplier.CID, "_id": ObjectId}, bson.M{"$set": updateSupplier})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": SupplierUpdated.MatchedCount})

	}
}

func DeleteSupplier() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		oid := c.GetHeader("oid")
		cid := c.GetHeader("cid")

		defer cancel()

		ObjectId, _ := primitive.ObjectIDFromHex(oid)
		deleteResult, err := supplierCollection.DeleteOne(ctx, bson.M{"cid": cid, "_id": ObjectId})
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
