package controllers

import (
	"promg/configs"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var clientCollection *mongo.Collection = configs.GetCollection(configs.DB, "clients")
var validateClient = validator.New()
