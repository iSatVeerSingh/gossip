package models

import (
	"time"

	"github.com/iSatVeerSingh/gossip/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageModel struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SentBy     primitive.ObjectID `json:"sentBy" bson:"sentBy"`
	ReceivedBy primitive.ObjectID `json:"receivedBy" bson:"receivedBy"`
	Time       time.Time          `json:"time" bson:"time"`
	MsgText    string             `json:"msgText" bson:"msgText"`
	Type       string             `json:"type" bson:"type"`
	IsRead     bool               `json:"isRead" bson:"isRead"`
}

type ParticipantUser struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
}

type ChatModel struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Users    [2]ParticipantUser `json:"users" bson:"users,omitempty"`
	Created  time.Time          `json:"created" bson:"created"`
	Type     string             `json:"type" bson:"type"`
	Messages []MessageModel     `json:"messages" bson:"messages"`
}

func GetChatCollection() *mongo.Collection {
	return db.GetMongoCollection("chats")
}
