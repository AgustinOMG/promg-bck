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

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var companyCollection *mongo.Collection = configs.GetCollection(configs.DB, "company")
var validateUser = validator.New()
var validateCompany = validator.New()

// *********************************************************      USERS
func NewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newUser models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateUser.Struct(&newUser); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newUserCreated, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		//return Success At creation
		c.JSON(http.StatusCreated, responses.PMGResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": newUserCreated}})
	}
}

func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		uid := c.GetHeader("uid")
		var user models.User
		defer cancel()

		err := userCollection.FindOne(ctx, bson.M{"uid": uid}).Decode(&user)
		if err != nil {
			println(err.Error())
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

func CheckUserAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		email := c.GetHeader("email")
		var user models.User
		defer cancel()

		err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"status": "none"}})
			return
		}
		if user.Status == "registered" {
			c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"status": "registered"}})
		}
		if user.Status == "invited" {
			c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"status": "invited"}})
		}

	}
}

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		uid := c.GetHeader("uid")
		email := c.GetHeader("email")
		println(uid)
		println(email)
		//var user models.User
		defer cancel()
		userRegistered, err := userCollection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"status": "registered", "uid": uid}})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"status": "none"}})
			return
		}
		c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": userRegistered}})
	}
}

func GetAllStaff() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		companyId := c.GetHeader("companyId")
		var staff []models.User
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{"cid": companyId})

		if err = results.All(context.TODO(), &staff); err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": staff}})

	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.GetHeader("userId")
		var updateUser models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&updateUser); err != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateUser.Struct(&updateUser); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		userUpdated, err := userCollection.UpdateOne(ctx, bson.M{"uid": userId}, bson.M{"$set": updateUser})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": userUpdated}})

	}
}

func DeleteAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		userDeleted, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if userDeleted.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.PMGResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": 0}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": userDeleted.DeletedCount}},
		)
	}
}

// *********************************************************      Company

func GetACompany() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		objId, _ := primitive.ObjectIDFromHex(cid)
		var company models.Company
		defer cancel()
		//TODO necesito cambair la variable de resultado asia generico, extraer el OBject ID y luego pasar esa info al modelos de company y regresar eso
		err := companyCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&company)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": company}})
	}
}

func UpdateCompany() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		uid := c.GetHeader("uid")
		var companyData models.Company
		defer cancel()

		//validate the request body
		if err := c.ShouldBindJSON(&companyData); err != nil {
			println("holaa")
			println(err.Error())
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateCompany.Struct(&companyData); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.PMGResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		objId, _ := primitive.ObjectIDFromHex(companyData.Cid)
		// Update the company data
		companyUpdated, err := companyCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": companyData})

		// if there is no company with this ID , a new company is created

		if companyUpdated.MatchedCount == 0 {
			// trasnefer the information from the json to a new variable avoid the cid , as is not needed
			newCompany := models.NewCompany{
				Name:   companyData.Name,
				Rfc:    companyData.Rfc,
				Street: companyData.Street,
				City:   companyData.City,
				State:  companyData.State,
				PC:     companyData.PC,
				Logo:   companyData.Logo,
				Conf:   companyData.Conf,
			}

			// insert a new company
			newCompanyCreated, err := companyCollection.InsertOne(ctx, newCompany)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			comId := newCompanyCreated.InsertedID.(primitive.ObjectID).Hex()

			_, erru := userCollection.UpdateOne(ctx, bson.M{"uid": uid}, bson.M{"$set": bson.M{"company.0": comId}})

			if erru != nil {
				c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": erru.Error()}})
				return
			}

			c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": true}})

		} else {

			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.PMGResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			// compid := companyUpdated.UpsertedID.(primitive.ObjectID).Hex()
			c.JSON(http.StatusOK, responses.PMGResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": true}})
		}

	}
}
