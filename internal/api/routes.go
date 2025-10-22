package api

import (
	"net/http"

	"github.com/andreposman/auction-house-api/internal/jsonutils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *API) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger, api.Sessions.LoadAndSave)
	api.Router.Route("/api", func(r chi.Router) {

		r.Get("/ping", api.pingRoute)

		r.Route("/v1", func(r chi.Router) {

			r.Route("/users", func(r chi.Router) {
				r.Post("/signup", api.handleSignupUser)
				r.Post("/login", api.handleLoginUser)
				r.With(api.AuthMiddleware).Post("/logout", api.handleLogoutUser)
			})
		})
	})
}

func (api *API) pingRoute(w http.ResponseWriter, r *http.Request) {
	_ = jsonutils.Encode(w, r, http.StatusOK, map[string]any{
		"message": "pong"})
}
