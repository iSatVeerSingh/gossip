package models

import (
	"time"

	"github.com/iSatVeerSingh/gossip/db"
	"github.com/iSatVeerSingh/gossip/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	Id          primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string              `json:"name" bson:"name"`
	Email       string              `json:"email" bson:"email"`
	Username    string              `json:"username" bson:"username"`
	Password    string              `json:"password" bson:"password"`
	Avatar      string              `json:"avatar" bson:"avatar"`
	About       string              `json:"about" bson:"about"`
	Status      string              `json:"status" bson:"status"`
	Created     time.Time           `json:"created" bson:"created"`
	Updated     time.Time           `json:"updated" bson:"updated"`
	Requests    []utils.RequestUser `json:"requests" bson:"requests"`
	Connections []utils.Connection  `json:"connections" bson:"connections"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUserCollection() *mongo.Collection {
	return db.GetMongoCollection("users")
}
