package utils

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type RequestUser struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Username    string             `json:"username" bson:"username"`
	RequestTime time.Time          `json:"requestTime,omitempty" bson:"requestTime,omitempty"`
}

type Connection struct {
	UserId         primitive.ObjectID `json:"userId" bson:"userId"`
	ConnectionTime time.Time          `json:"connectionTime" bson:"connectionTime"`
	ConnectionType string             `json:"connectionType" bson:"connectionType"`
}

type ConnectionRequest struct {
	RequestedUser RequestUser `json:"requestedUser"`
	RequestedBy   RequestUser `json:"requestedBy"`
}

type AcceptRequest struct {
	AcceptedUser RequestUser `json:"accepteduser"`
	AcceptedBy   RequestUser `json:"acceptedby"`
}

type User struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	Username    string             `json:"username,omitempty" bson:"username,omitempty"`
	Avatar      string             `json:"avatar,omitempty" bson:"avatar,omitempty"`
	About       string             `json:"about,omitempty" bson:"about,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	Requests    []RequestUser      `json:"requests,omitempty" bson:"requests,omitempty"`
	Connections []RequestUser      `json:"connections,omitempty" bson:"connections,omitempty"`
}

type CtxUserInfoKey struct{}
