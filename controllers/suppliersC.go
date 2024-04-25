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

var supplierCollection *mongo.Collection = configs.GetCollection(configs.DB, "suppliers")
var validateSupplier = validator.New()

func NewSupplier() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newSupplier models.Supplier
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&newSupplier); err != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateSupplier.Struct(&newSupplier); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newSupplierCreated, err := clientCollection.InsertOne(ctx, newSupplier)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		//return Success At creation
		c.JSON(http.StatusCreated, responses.PMGResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": newSupplierCreated}})
	}
}

func GetSuppliers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		var suppliers []models.Supplier
		defer cancel()

		results, err := supplierCollection.Find(ctx, bson.M{"cid": cid})

		defer results.Close(ctx)
		if err = results.All(context.TODO(), &suppliers); err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": suppliers}})

	}
}

func UpdateSupplier() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		var updateSupplier models.Supplier
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&updateSupplier); err != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateSupplier.Struct(&updateSupplier); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		supplierUpdated, err := supplierCollection.UpdateOne(ctx, bson.M{"cid": cid}, bson.M{"$set": updateSupplier})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": supplierUpdated}})

	}
}

func DeleteSupplier() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.Param("cid")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(cid)

		supplierDeleted, err := supplierCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if supplierDeleted.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.PMGResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": 0}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": supplierDeleted.DeletedCount}},
		)
	}
}
