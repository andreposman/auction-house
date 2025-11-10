package api

import (
	"net/http"

	"github.com/andreposman/auction-house-api/internal/jsonutils"
	"github.com/gorilla/csrf"
)

func (api *API) HandleGetCSRFToken(w http.ResponseWriter, r *http.Request) {
	token := csrf.Token(r)
	jsonutils.Encode(w, r, http.StatusOK, map[string]any{
		"csrf_token": token,
	})

}

func (api *API) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !api.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			jsonutils.Encode(w, r, http.StatusUnauthorized, map[string]any{
				"message": "must be logged in",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}
