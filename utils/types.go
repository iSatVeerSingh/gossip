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

type AuthUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstname" bson:"firstname"`
	LastName  string             `json:"lastname" bson:"lastname"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Avatar    string             `json:"avatar" bson:"avatar"`
	About     string             `json:"about" bson:"about"`
	Status    string             `json:"status,omitempty" bson:"status,omitempty"`
}

type CtxUserInfoKey struct{}
