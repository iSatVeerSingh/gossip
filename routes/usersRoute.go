package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iSatVeerSingh/gossip/controllers"
)

func SetUsersRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc(USER_REGISTER, controllers.CreateUser).Methods(http.MethodPost)

	return router
}
