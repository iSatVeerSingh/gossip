package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

type RequestUser struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
}

type ConnectionRequest struct {
	RequestedTo RequestUser `json:"requestedto"`
	RequestedBy RequestUser `json:"requestedby"`
}

type AcceptRequest struct {
	AcceptedUser RequestUser `json:"accepteduser"`
	AcceptedBy   RequestUser `json:"acceptedby"`
}

type AuthUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type User struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName   string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName    string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	Username    string             `json:"username,omitempty" bson:"username,omitempty"`
	Avatar      string             `json:"avatar,omitempty" bson:"avatar,omitempty"`
	About       string             `json:"about,omitempty" bson:"about,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	Requests    []RequestUser      `json:"requests,omitempty" bson:"requests,omitempty"`
	Connections []RequestUser      `json:"connections,omitempty" bson:"connections,omitempty"`
}

type CtxUserInfoKey struct{}
