package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/iSatVeerSingh/gossip/models"
	"github.com/iSatVeerSingh/gossip/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddRequest(request utils.ConnectionRequest) error {

	countFilter := bson.M{"_id": request.RequestedTo.Id, "requests": bson.M{
		"$elemMatch": bson.M{"_id": request.RequestedBy.Id, "username": request.RequestedBy.Username},
	}}

	result, err := models.GetUserCollection().CountDocuments(context.TODO(), countFilter)

	if result != 0 {
		return errors.New("request already sent")
	}

	if err != nil {
		return err
	}

	updateFilterQuery := bson.M{"_id": request.RequestedTo.Id}

	updateRequestData := bson.M{"$push": bson.M{
		"requests": request.RequestedBy,
	}}

	updateresult, err := models.GetUserCollection().UpdateOne(context.TODO(), updateFilterQuery, updateRequestData)

	if err != nil {
		return err
	}

	if updateresult.ModifiedCount == 0 {
		return errors.New("internal server error")
	}
	return nil
}

func GetAllRequestsByUser(id string) (utils.User, error) {
	var requests utils.User

	mongoId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return requests, err
	}

	filter := bson.M{"_id": mongoId}

	projectFilter := bson.M{"requests": 1}

	result := models.GetUserCollection().FindOne(context.TODO(), filter, options.FindOne().SetProjection(projectFilter))

	fmt.Println(result)

	err = result.Decode(&requests)

	if err != nil {
		return requests, err
	}

	return requests, nil
}

func CreateConversation(request utils.AcceptRequest) (interface{}, error) {
	countFilter := bson.M{"_id": request.AcceptedBy.Id, "requests": bson.M{
		"$elemMatch": bson.M{
			"_id":      request.AcceptedUser.Id,
			"username": request.AcceptedUser.Username,
		},
	}}

	result, err := models.GetUserCollection().CountDocuments(context.TODO(), countFilter)

	if result == 0 {
		return "", errors.New("user does not exists")
	}

	if err != nil {
		return "", err
	}

	updateFilterQuery := bson.M{"_id": request.AcceptedBy.Id}

	updateRequestData := bson.M{
		"$pull": bson.M{"requests": request.AcceptedUser},
		"$push": bson.M{"connections": request.AcceptedUser},
	}

	updateResult, err := models.GetUserCollection().UpdateOne(context.TODO(), updateFilterQuery, updateRequestData)

	if updateResult.ModifiedCount == 0 {
		return "", errors.New("internal server error")
	}

	if err != nil {
		return "", err
	}

	chatResult, err := CreateChat(request)

	return chatResult, err
}

func CreateChat(users utils.AcceptRequest) (*mongo.InsertOneResult, error) {
	var chat models.ChatModel

	chat.Users[0].Id = users.AcceptedBy.Id
	chat.Users[0].Username = users.AcceptedBy.Username

	chat.Users[1].Id = users.AcceptedUser.Id
	chat.Users[1].Username = users.AcceptedUser.Username

	chat.Type = "one-to-one"

	chat.Created = primitive.NewDateTimeFromTime(time.Now())

	result, err := models.GetChatCollection().InsertOne(context.TODO(), chat)

	return result, err
}

func GetAllConnectionsByUser(id string) (utils.User, error) {
	var connections utils.User

	mongoId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return connections, err
	}

	filter := bson.M{"_id": mongoId}

	projectFilter := bson.M{"connections": 1}

	result := models.GetUserCollection().FindOne(context.TODO(), filter, options.FindOne().SetProjection(projectFilter))

	fmt.Println(result)

	err = result.Decode(&connections)

	if err != nil {
		return connections, err
	}

	return connections, nil
}
