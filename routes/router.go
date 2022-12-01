package routes

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter = SetAuthRoutes(apiRouter)

	apiRouter = SetUserRoutes(apiRouter)

	return apiRouter
}
