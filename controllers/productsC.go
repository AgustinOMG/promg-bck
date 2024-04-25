package controllers

import (
	"context"
	"net/http"
	"promg/configs"
	"promg/models"
	"promg/responses"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")
var validateproduct = validator.New()

func Newproduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newproduct models.Product
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&newproduct); err != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateproduct.Struct(&newproduct); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newproductCreated, err := productCollection.InsertOne(ctx, newproduct)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		//return Success At creation
		c.JSON(http.StatusCreated, responses.PMGResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": newproductCreated}})
	}
}

func Getproducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		var products []models.Product
		defer cancel()

		results, err := productCollection.Find(ctx, bson.M{"cid": cid})

		defer results.Close(ctx)
		if err = results.All(context.TODO(), &products); err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": products}})

	}
}

func Updateproduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		var updateproduct models.Product
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&updateproduct); err != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateproduct.Struct(&updateproduct); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		productUpdated, err := productCollection.UpdateOne(ctx, bson.M{"cid": cid}, bson.M{"$set": updateproduct})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": productUpdated}})

	}
}

func DeleteAproduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.Param("cid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(cid)

		productDeleted, err := productCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if productDeleted.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.PMGResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": 0}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": productDeleted.DeletedCount}},
		)
	}
}
