package controllers

import (
	"promg/configs"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var quotesCollection *mongo.Collection = configs.GetCollection(configs.DB, "quotes")
var validateQuotes = validator.New()
