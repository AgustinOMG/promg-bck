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

var itemCollection *mongo.Collection = configs.GetCollection(configs.DB, "items")
var validateitem = validator.New()

func Newitem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newitem models.Item
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&newitem); err != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateitem.Struct(&newitem); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newitemCreated, err := itemCollection.InsertOne(ctx, newitem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		//return Success At creation
		c.JSON(http.StatusCreated, responses.PMGResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": newitemCreated}})
	}
}

func Getitems() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		var items []models.Item
		defer cancel()

		results, err := itemCollection.Find(ctx, bson.M{"cid": cid})

		defer results.Close(ctx)
		if err = results.All(context.TODO(), &items); err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": items}})

	}
}

func Updateitem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		var updateitem models.Item
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&updateitem); err != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateitem.Struct(&updateitem); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		itemUpdated, err := itemCollection.UpdateOne(ctx, bson.M{"cid": cid}, bson.M{"$set": updateitem})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": itemUpdated}})

	}
}

func DeleteAitem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.Param("cid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(cid)

		itemDeleted, err := itemCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if itemDeleted.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.PMGResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": 0}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": itemDeleted.DeletedCount}},
		)
	}
}
