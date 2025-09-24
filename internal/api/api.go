package api

import (
	"net/http"

	"github.com/andreposman/action-house-api/internal/services"
	"github.com/go-chi/chi/v5"
)

type API struct {
	Router      *chi.Mux
	UserService services.UserService
}

func (api *API) handleCreateUser(w http.ResponseWriter, r *http.Request) {

}
