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

var clientCollection *mongo.Collection = configs.GetCollection(configs.DB, "clients")
var validateClient = validator.New()

func NewClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newClient models.Client
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&newClient); err != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateUser.Struct(&newClient); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newClientCreated, err := clientCollection.InsertOne(ctx, newClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		//return Success At creation
		c.JSON(http.StatusCreated, responses.PMGResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": newClientCreated}})
	}
}

func GetClients() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		var clients []models.Client
		defer cancel()

		results, err := clientCollection.Find(ctx, bson.M{"cid": cid})

		defer results.Close(ctx)
		if err = results.All(context.TODO(), &clients); err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": clients}})

	}
}

func UpdateClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		var updateClient models.Client
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&updateClient); err != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateUser.Struct(&updateClient); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		clientUpdated, err := userCollection.UpdateOne(ctx, bson.M{"cid": cid}, bson.M{"$set": updateClient})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": clientUpdated}})

	}
}

func DeleteAClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.Param("cid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(cid)

		clientDeleted, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if clientDeleted.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.PMGResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": 0}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": clientDeleted.DeletedCount}},
		)
	}
}
