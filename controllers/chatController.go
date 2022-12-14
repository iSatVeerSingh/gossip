package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/iSatVeerSingh/gossip/helpers"
	"github.com/iSatVeerSingh/gossip/services"
	"github.com/iSatVeerSingh/gossip/utils"
)

func AddRequest(w http.ResponseWriter, r *http.Request) {
	var request utils.ConnectionRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		helpers.GetErrorResponse(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	userInfo := r.Context().Value(utils.CtxUserInfoKey{}).(utils.AuthUser)

	if userInfo.Id != request.RequestedBy.Id.Hex() || userInfo.Id == request.RequestedUser.Id.Hex() {
		helpers.GetErrorResponse(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	err = services.AddNewRequest(request)

	if err != nil {
		helpers.GetErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.GetSuccessResponse(w, "Connection request sent successfully", http.StatusOK)

}

func GetAllRequests(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value(utils.CtxUserInfoKey{}).(utils.AuthUser)

	result, err := services.GetAllRequestsByUser(userInfo.Id)

	if err != nil {
		helpers.GetErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.GetSuccessResponse(w, result, http.StatusOK)

}

func CreateConnection(w http.ResponseWriter, r *http.Request) {
	var acceptRequest utils.AcceptRequest

	err := json.NewDecoder(r.Body).Decode(&acceptRequest)

	if err != nil {
		helpers.GetErrorResponse(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	userInfo := r.Context().Value(utils.CtxUserInfoKey{}).(utils.AuthUser)

	if userInfo.Id != acceptRequest.AcceptedBy.Id.Hex() || userInfo.Id == acceptRequest.AcceptedUser.Id.Hex() {
		helpers.GetErrorResponse(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	result, err := services.CreateConnection(acceptRequest)

	if err != nil {
		helpers.GetErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.GetSuccessResponse(w, result, http.StatusOK)
}

func GetAllConnections(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value(utils.CtxUserInfoKey{}).(utils.AuthUser)

	result, err := services.GetAllConnectionsByUser(userInfo.Id)

	if err != nil {
		helpers.GetErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.GetSuccessResponse(w, result, http.StatusOK)
}
