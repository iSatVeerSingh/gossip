package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageModel struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SentBy     primitive.ObjectID `json:"sentBy" bson:"sentBy"`
	ReceivedBy primitive.ObjectID `json:"receivedBy" bson:"receivedBy"`
	Time       time.Time          `json:"time" bson:"time"`
	MsgText    string             `json:"msgText" bson:"msgText"`
	Type       string             `json:"type" bson:"type"`
}

type ChatModel struct {
	Id       primitive.ObjectID    `json:"id" bson:"_id,omitempty"`
	Users    [2]primitive.ObjectID `json:"users" bson:"users"`
	Created  time.Time             `json:"created" bson:"created"`
	Type     string                `json:"type" bson:"type"`
	Messages []MessageModel        `json:"messages" bson:"messages"`
}
