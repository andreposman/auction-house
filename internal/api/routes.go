package api

import (
	"net/http"

	"github.com/andreposman/action-house-api/internal/jsonutils"
	"github.com/go-chi/chi/v5"
)

func (api *API) BindRoutes() {
	api.Router.Route("/api", func(r chi.Router) {

		r.Get("/ping", api.pingRoute)

		r.Route("/v1", func(r chi.Router) {

			r.Route("/users", func(r chi.Router) {
				r.Post("/signup", api.handleSignupUser)
				r.Post("/login", api.handleLoginUser)
				r.Post("/logout", api.handleLogoutUser)
			})
		})
	})
}

func (api *API) pingRoute(w http.ResponseWriter, r *http.Request) {
	_ = jsonutils.Encode(w, r, http.StatusUnprocessableEntity, map[string]any{
		"message": "pong"})
}
