package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andreposman/action-house-api/internal/jsonutils"
	"github.com/andreposman/action-house-api/internal/services"
	"github.com/andreposman/action-house-api/internal/usecase/user"
)

func (api *API) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.CreateUserReq](r)

	fmt.Printf("\n\nUsername(Len: %v): %v \n", len(data.UserName), data.UserName)
	fmt.Printf("Email(Len: %v): %v \n", len(data.Email), data.Email)
	fmt.Printf("Bio(Len: %v): %v \n", len(data.Bio), data.Bio)
	fmt.Printf("Password(Len: %v): %v \n", len(data.Password), data.Password)

	if err != nil {
		_ = jsonutils.Encode(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	fmt.Printf("\nRequest: %v \n", data)

	id, err := api.UserService.CreateUser(r.Context(),
		data.UserName,
		data.Email,
		data.Password,
		data.Bio)
	if err != nil {
		if errors.Is(err, services.ErrDuplicateUniqueField) {
			_ = jsonutils.Encode(w, r, http.StatusUnprocessableEntity, map[string]any{
				"error": "email or username already exists",
			})
			return
		}

		_ = jsonutils.Encode(w, r, http.StatusUnprocessableEntity, map[string]any{
			"user_id": id,
		})
		return
	}

	_ = jsonutils.Encode(w, r, http.StatusOK, map[string]any{
		"user_id": id,
	})
}

func (api *API) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	panic("TODO")
}

func (api *API) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	panic("TODO")
}
