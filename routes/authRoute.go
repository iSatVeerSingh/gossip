package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iSatVeerSingh/gossip/controllers"
	"github.com/iSatVeerSingh/gossip/middlewares"
)

// User authentication routes
func SetAuthRoutes(router *mux.Router) *mux.Router {
	authRouter := mux.NewRouter()

	authRouter.HandleFunc(USER_REGISTER, controllers.CreateUser).Methods(http.MethodPost)
	authRouter.HandleFunc(USER_LOGIN, controllers.LoginUser).Methods(http.MethodPost)
	authRouter.HandleFunc(USER_STATUS, controllers.LoginStatus).Methods(http.MethodGet)
	authRouter.HandleFunc(USER_LOGOUT, controllers.LogoutUser).Methods(http.MethodPost)

	router.PathPrefix("/auth").Handler(middlewares.Authorize(authRouter))

	return router
}
