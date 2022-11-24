package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/iSatVeerSingh/gossip/helpers"
	"github.com/iSatVeerSingh/gossip/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user *models.User) (interface{}, error) {
	hashPassword, ok := helpers.EncryptPassword(user.Password)

	if !ok {
		return "", errors.New("internal server error")
	}

	user.Password = hashPassword

	result, err := models.GetUserCollection().InsertOne(context.TODO(), user)

	fmt.Println(err.(mongo.WriteException).WriteErrors[0].Details.Elements())
	return result, err
}
