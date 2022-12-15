package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iSatVeerSingh/gossip/helpers"
	"github.com/iSatVeerSingh/gossip/models"
	"github.com/iSatVeerSingh/gossip/services"
	"github.com/iSatVeerSingh/gossip/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserModel

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		helpers.GetErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if userError, isValid := helpers.CreateUserValidation(&user); !isValid {
		helpers.GetErrorResponse(w, userError, http.StatusBadRequest)
		return
	}

	result, err := services.CreateUser(&user)
	if err != nil {
		helpers.GetErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.GetSuccessResponse(w, result, http.StatusCreated)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.LoginUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.GetErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if loginError, isValid := helpers.LoginUserValidation(&user); !isValid {
		helpers.GetErrorResponse(w, loginError, http.StatusBadRequest)
		return
	}

	result, err := services.LoginUser(&user)
	if err != nil {
		helpers.GetErrorResponse(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token := helpers.GenerateToken(result)

	tokenCookie := http.Cookie{
		Name:   "Token",
		Value:  token,
		MaxAge: 60 * 60 * 24,
		// Secure: true,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, &tokenCookie)

	helpers.GetSuccessResponse(w, result, http.StatusOK)
}

func UserProfile(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	user, err := services.FindUserByUsername(username)

	if err == mongo.ErrNoDocuments {
		helpers.GetErrorResponse(w, "Couldn't find any user", http.StatusBadRequest)
		return
	}

	if err != nil {
		helpers.GetErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	helpers.GetSuccessResponse(w, user, http.StatusOK)
}

func LoginStatus(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value(utils.CtxUserInfoKey{}).(utils.AuthUser)

	user, err := services.FindUserByUsername(userInfo.Username)
	if err != nil {
		helpers.GetErrorResponse(w, "You are not authrized", http.StatusUnauthorized)
		return
	}

	helpers.GetSuccessResponse(w, user, http.StatusOK)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	tokenCookie := http.Cookie{
		Name:   "Token",
		Value:  "",
		MaxAge: -1,
		// Secure: true,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &tokenCookie)

	helpers.GetSuccessResponse(w, "Loggedout successfully", http.StatusOK)
}
