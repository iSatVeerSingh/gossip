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

	if userInfo.Id != request.RequestedBy.Id.Hex() || userInfo.Id == request.RequestedTo.Id.Hex() {
		helpers.GetErrorResponse(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	err = services.AddRequest(request)

	if err != nil {
		helpers.GetErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.GetSuccessResponse(w, "Connection request sent successfully", http.StatusOK)

}
