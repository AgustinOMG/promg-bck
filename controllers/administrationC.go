package controllers

import (
	"context"
	"fmt"
	"net/http"
	"promg/configs"
	"promg/models"
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

// *********************************************************      Company

func GetACompany() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		objId, _ := primitive.ObjectIDFromHex(cid)
		var company models.Company
		var result models.CompanyDb
		defer cancel()

		err := companyCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}

		company = models.Company(result)
		println(company.Conf.QFolio)

		c.JSON(http.StatusOK, gin.H{"data": company})
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
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateCompany.Struct(&companyData); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": validationErr.Error()})
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
				c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
				return
			}
			comId := newCompanyCreated.InsertedID.(primitive.ObjectID).Hex()

			_, erru := userCollection.UpdateOne(ctx, bson.M{"uid": uid}, bson.M{"$set": bson.M{"company.0": comId}})

			if erru != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"data": erru.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": true})

		} else {

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
				return
			}

			// compid := companyUpdated.UpsertedID.(primitive.ObjectID).Hex()
			c.JSON(http.StatusOK, gin.H{"data": true})
		}

	}
}

// *********************************************************      USERS
func NewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newUser models.User
		defer cancel()

		//validate the request body
		if err := c.ShouldBindJSON(&newUser); err != nil {
			println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateUser.Struct(&newUser); validationErr != nil {
			println(validationErr.Error())
			c.JSON(http.StatusBadRequest, gin.H{"data": validationErr.Error()})
			return
		}

		newUserCreated, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		//return Success At creation
		c.JSON(http.StatusCreated, gin.H{"data": newUserCreated})
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
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
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
			c.JSON(http.StatusInternalServerError, gin.H{"data": "none"})
			return
		}
		if user.Status == "registered" {
			c.JSON(http.StatusOK, gin.H{"data": "registered"})
		}
		if user.Status == "invited" {
			c.JSON(http.StatusOK, gin.H{"data": "invited"})
		}

	}
}

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		uid := c.GetHeader("uid")
		email := c.GetHeader("email")

		defer cancel()
		userRegistered, err := userCollection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"status": "registered", "uid": uid}})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": "none"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": userRegistered.MatchedCount})
	}
}

func GetStaff() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cid := c.GetHeader("cid")
		var staff []models.User
		defer cancel()
		filter := bson.D{{"company.0", cid}}

		results, err := userCollection.Find(ctx, filter)

		for results.Next(context.TODO()) {
			var elem models.User

			err := results.Decode(&elem)
			if err != nil {
				fmt.Println(err)
			}
			staff = append(staff, elem)
		}

		if err = results.All(context.TODO(), &staff); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": staff})

	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var updateUser models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&updateUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validateUser.Struct(&updateUser); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": validationErr.Error()})
			return
		}

		userUpdated, err := userCollection.UpdateOne(ctx, bson.M{"uid": updateUser.UID}, bson.M{"$set": updateUser})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": userUpdated.MatchedCount})

	}
}

/*
func DeleteAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		uid := c.Param("uid")
		defer cancel()

		app, errFire := firebase.NewApp(context.Background(), nil)
		if errFire != nil {
			log.Fatalf("error initializing app: %v\n", errFire)
		}

		client, errAuth := app.Auth(ctx)
		if errAuth != nil {
			println("error Authenticating: %v\n", errAuth.Error())

		}

		errDelete := client.DeleteUser(ctx, uid)
		if errDelete != nil {
			println("error deleting user: %v\n", errDelete.Error())
		}
		println("Successfully deleted user: %s\n", uid)

		userDeleted, err := userCollection.DeleteOne(ctx, bson.M{"uid": uid})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}

		if userDeleted.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				gin.H{"data": 0},
			)
			return
		}

		c.JSON(http.StatusOK,
			gin.H{"data": userDeleted.DeletedCount},
		)
	}
}
*/
// *********************************************************      Configuration

func UpdateConfiguration() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var updateConf models.Conf
		cid := c.GetHeader("cid")

		defer cancel()

		//validate the request body
		if err := c.ShouldBindJSON(&updateConf); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
			return
		}

		confUpdated, err := companyCollection.UpdateOne(ctx, bson.M{"cid": cid}, bson.M{"$set": bson.M{"conf": updateConf}})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": confUpdated.MatchedCount})

	}
}
