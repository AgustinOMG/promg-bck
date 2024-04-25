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
	"go.mongodb.org/mongo-driver/mongo"
)

var quotesCollection *mongo.Collection = configs.GetCollection(configs.DB, "quotes")
var validateQuotes = validator.New()

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
		if validationErr := validateQuotes.Struct(&quote); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		result, err := quotesCollection.InsertOne(ctx, quote)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.PMGResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

// ** updateQuote
//** getQuote
//** searhcQuote
//** deleteQuote
//**searchQuoteClient
