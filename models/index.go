package models

import (
	"context"
	"log"

	"github.com/iSatVeerSingh/gossip/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddIndexes() {
	var err error

	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	usernameIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err = db.GetMongoCollection(USER_COLL).Indexes().CreateMany(context.TODO(), []mongo.IndexModel{emailIndex, usernameIndex})
	if err != nil {
		log.Fatal(err)
	}
}
