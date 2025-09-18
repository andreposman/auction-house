package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type API struct {
	Router *chi.Mux
}

func (api *API) handleCreateUser(w http.ResponseWriter, r *http.Request) {

}
