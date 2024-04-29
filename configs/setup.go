package configs

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	connString := getConnString()
	clientOpts := options.Client().ApplyURI(
		connString)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {

		log.Fatal(err)
	}
	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("promg").Collection(collectionName)
	return collection
}
