package services

import (
	"context"
	"errors"

	"github.com/iSatVeerSingh/gossip/models"
	"github.com/iSatVeerSingh/gossip/utils"
	"go.mongodb.org/mongo-driver/bson"
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
