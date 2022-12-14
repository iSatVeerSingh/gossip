package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Send ok for preflight request by browser
	router.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router = SetAuthRoutes(router)

	router = SetUserRoutes(router)

	return router
}
