package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/iSatVeerSingh/gossip/helpers"
	"github.com/iSatVeerSingh/gossip/models"
	"github.com/iSatVeerSingh/gossip/services"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

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
		MaxAge: int(time.Minute) * 15,
		// Secure: true,
		HttpOnly: true,
	}

	http.SetCookie(w, &tokenCookie)

	helpers.GetSuccessResponse(w, result, http.StatusOK)
}
