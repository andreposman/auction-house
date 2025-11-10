package api

import (
	"net/http"
	"os"

	"github.com/andreposman/auction-house-api/internal/jsonutils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
)

func (api *API) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger, api.Sessions.LoadAndSave)

	csrfMiddleware := csrf.Protect(
		[]byte(os.Getenv("AUCTION_HOUSE_CSRF_KEY")),
		csrf.Secure(false), //dev only, set http instead https, later add to .env
		csrf.TrustedOrigins([]string{"http://localhost:3080"}),
		csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jsonutils.Encode(w, r, http.StatusForbidden, map[string]any{
				"error": "CSRF token validation failed",
			})
		})))

	api.Router.Use(csrfMiddleware)

	api.Router.Route("/api", func(r chi.Router) {

		r.Get("/ping", api.pingRoute)

		r.Route("/v1", func(r chi.Router) {

			r.Get("/csrf", api.HandleGetCSRFToken)

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
