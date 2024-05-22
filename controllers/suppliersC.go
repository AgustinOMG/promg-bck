package controllers

import (
	"promg/configs"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var supplierCollection *mongo.Collection = configs.GetCollection(configs.DB, "suppliers")
var validateSupplier = validator.New()
