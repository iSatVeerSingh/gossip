package services

import (
	"context"
	"errors"
	"time"

	"github.com/iSatVeerSingh/gossip/helpers"
	"github.com/iSatVeerSingh/gossip/models"
	"github.com/iSatVeerSingh/gossip/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateNewUser(user *models.UserModel) (interface{}, error) {
	hashPassword, ok := helpers.EncryptPassword(user.Password)

	if !ok {
		return "", errors.New("internal server error")
	}

	user.Password = hashPassword
	user.Created = time.Now()
	user.Updated = time.Now()
	user.Requests = make([]utils.RequestUser, 0)
	user.Connections = make([]utils.Connection, 0)

	result, err := models.GetUserCollection().InsertOne(context.TODO(), user)

	if mongo.IsDuplicateKeyError(err) {
		return nil, err
	}
	return result, err
}

func LoginUser(loginData *models.LoginUser) (utils.AuthUser, error) {
	user, err := GetUser("username", loginData.Username)

	var authUser utils.AuthUser

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

func GetUser(key string, value string) (models.UserModel, error) {
	var user models.UserModel

	result := models.GetUserCollection().FindOne(context.TODO(), bson.D{{Key: key, Value: value}})

	err := result.Decode(&user)
	return user, err
}

func FindUserByUsername(username string) (utils.User, error) {
	var user utils.User

	projectFilter := bson.M{
		"name":     1,
		"username": 1,
		"avatar":   1,
		"about":    1,
		"status":   1,
	}

	result := models.GetUserCollection().FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}, options.FindOne().SetProjection(projectFilter))

	err := result.Decode(&user)

	return user, err
}
