package db

import (
	"context"
	"log"

	"github.com/iSatVeerSingh/gossip/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func GetMongoSession() *mongo.Client {
	var err error
	mongoUri := utils.GetEnv("MONGODB_URI")

	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))

	if err != nil {
		log.Fatal(err)
	}

	return mongoClient
}

func GetMongoDatabase() *mongo.Database {
	dbname := utils.GetEnv("DB_NAME")

	database := mongoClient.Database(dbname)
	return database
}

func GetMongoCollection(collName string) *mongo.Collection {
	coll := GetMongoDatabase().Collection(collName)

	return coll
}

func MongoSessionClose(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
