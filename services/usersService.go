package services

import (
	"context"
	"errors"

	"github.com/iSatVeerSingh/gossip/helpers"
	"github.com/iSatVeerSingh/gossip/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user *models.User) (interface{}, error) {
	hashPassword, ok := helpers.EncryptPassword(user.Password)

	if !ok {
		return "", errors.New("internal server error")
	}

	user.Password = hashPassword

	result, err := models.GetUserCollection().InsertOne(context.TODO(), user)

	if mongo.IsDuplicateKeyError(err) {
		return nil, errors.New("user already exists")
	}
	return result, err
}

func LoginUser(loginData *models.LoginUser) (AuthUser, error) {
	user, err := GetUserByUsername(loginData.Username)

	var authUser AuthUser

	if err != nil {
		return authUser, errors.New("invalid credentials")
	}

	if isValid := helpers.ValidatePassword(loginData.Password, user.Password); !isValid {
		return authUser, errors.New("invalid credentials")
	}

	authUser.Id = user.Id.Hex()
	authUser.Username = user.Username

	return authUser, nil
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	result := models.GetUserCollection().FindOne(context.TODO(), bson.D{{Key: "username", Value: username}})
	err := result.Decode(&user)
	if err == mongo.ErrNoDocuments {
		return user, errors.New("couldn't find any user")
	}
	return user, err
}
