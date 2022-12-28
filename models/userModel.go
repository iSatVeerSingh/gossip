package models

import (
	"time"

	"github.com/iSatVeerSingh/gossip/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Email       string             `json:"email" bson:"email"`
	Username    string             `json:"username" bson:"username"`
	Password    string             `json:"password" bson:"password"`
	Avatar      string             `json:"avatar" bson:"avatar"`
	About       string             `json:"about" bson:"about"`
	Status      string             `json:"status" bson:"status"`
	Created     time.Time          `json:"created" bson:"created"`
	Updated     time.Time          `json:"updated" bson:"updated"`
	Requests    []RequestUser      `json:"requests" bson:"requests"`
	Connections []Connection       `json:"connections" bson:"connections"`
}

type RequestUser struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Username    string             `json:"username" bson:"username"`
	RequestTime time.Time          `json:"requestTime" bson:"requestTime"`
}

type Connection struct {
	UserId         primitive.ObjectID `json:"userId" bson:"userId"`
	ConnectionTime time.Time          `json:"connectionTime" bson:"connectionTime"`
	ConnectionType string             `json:"connectionType" bson:"connectionType"`
}

// type UserModel struct {
// 	Id          primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
// 	FirstName   string              `json:"firstname" bson:"firstname"`
// 	LastName    string              `json:"lastname" bson:"lastname"`
// 	Email       string              `json:"email" bson:"email"`
// 	Username    string              `json:"username" bson:"username"`
// 	Password    string              `json:"password" bson:"password"`
// 	Avatar      string              `json:"avatar" bson:"avatar"`
// 	About       string              `json:"about" bson:"about"`
// 	Status      string              `json:"status" bson:"status"`
// 	Created     time.Time           `json:"created" bson:"created"`
// 	Updated     time.Time           `json:"updated" bson:"updated"`
// 	Requests    []utils.RequestUser `json:"requests" bson:"requests,omitempty"`
// 	Connections []string            `json:"connections" bson:"connections,omitempty"`
// }

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUserCollection() *mongo.Collection {
	return db.GetMongoCollection("users")
}
