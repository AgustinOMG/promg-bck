package controllers

import (
	"promg/configs"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")
var validateproduct = validator.New()
