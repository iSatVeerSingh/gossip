package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iSatVeerSingh/gossip/middlewares"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	userRouter := mux.NewRouter()

	userRouter.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is user router"))
	})

	router.PathPrefix("/users").Handler(middlewares.Authorize(userRouter))

	return router
}
