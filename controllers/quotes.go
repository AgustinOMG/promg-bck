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

var companyCollection *mongo.Collection = configs.GetCollection(configs.DB, "company")
var validateCompany = validator.New()

func NewQuote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var quote models.Quote
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&quote); err != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateCompany.Struct(&quote); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		result, err := companyCollection.InsertOne(ctx, quote)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.PMGResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetACompany() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		CompanyId := c.GetHeader("CompanyId")
		var Company models.Company
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(CompanyId)

		err := companyCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&Company)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": Company}})
	}
}

// ** updateQuote
//** getQuote
//** searhcQuote
//** deleteQuote
//**searchQuoteClient
