package middlewares

import (
	"net/http"

	"github.com/iSatVeerSingh/gossip/helpers"
)

func Authorize(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("Token")

		if err != nil {
			helpers.GetErrorResponse(w, "You are not authorized", http.StatusUnauthorized)
			return
		}

		_, ok := helpers.ValidateToken(tokenCookie.Value)
		if !ok {
			helpers.GetErrorResponse(w, "You are not authorized", http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
