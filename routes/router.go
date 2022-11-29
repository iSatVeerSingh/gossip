package routes

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router = SetAuthRoutes(router)

	router = SetUserRoutes(router)

	return router
}
