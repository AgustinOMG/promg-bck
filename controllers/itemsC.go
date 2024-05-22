package controllers

import (
	"promg/configs"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var itemCollection *mongo.Collection = configs.GetCollection(configs.DB, "items")
var validateitem = validator.New()
