package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iSatVeerSingh/gossip/controllers"
	"github.com/iSatVeerSingh/gossip/middlewares"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	userRouter := mux.NewRouter()

	userRouter.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is user router"))
	})

	userRouter.HandleFunc(USERS_REQUEST, controllers.GetAllRequests).Methods(http.MethodGet)
	userRouter.HandleFunc(USERS_REQUEST, controllers.AddRequest).Methods(http.MethodPatch)
	userRouter.HandleFunc(USERS_REQUEST, controllers.CreateConversation).Methods(http.MethodPost)
	userRouter.HandleFunc(USERS_CONNECTIONS, controllers.GetAllConnections).Methods(http.MethodGet)
	userRouter.HandleFunc(USER_PROFILE, controllers.UserProfile).Methods(http.MethodGet)

	router.PathPrefix("/users").Handler(middlewares.Authorize(userRouter))

	return router
}
