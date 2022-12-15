package middlewares

import (
	"context"
	"net/http"

	"github.com/iSatVeerSingh/gossip/helpers"
	"github.com/iSatVeerSingh/gossip/utils"
)

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path == "/auth/register" || path == "/auth/login" {
			next.ServeHTTP(w, r)
			return
		}

		tokenCookie, err := r.Cookie("Token")

		if err != nil {
			helpers.GetErrorResponse(w, "You are not authorized", http.StatusUnauthorized)
			return
		}

		userInfo, ok := helpers.ValidateToken(tokenCookie.Value)
		if !ok {
			helpers.GetErrorResponse(w, "You are not authorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), utils.CtxUserInfoKey{}, userInfo)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
