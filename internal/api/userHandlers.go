package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andreposman/auction-house-api/internal/jsonutils"
	"github.com/andreposman/auction-house-api/internal/services"
	"github.com/andreposman/auction-house-api/internal/usecase/user"
)

func (api *API) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.CreateUserReq](r)
	if err != nil {
		_ = jsonutils.Encode(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	fmt.Printf("\n\nUsername(Len: %v): %v \n", len(data.UserName), data.UserName)
	fmt.Printf("Email(Len: %v): %v \n", len(data.Email), data.Email)
	fmt.Printf("Bio(Len: %v): %v \n", len(data.Bio), data.Bio)
	fmt.Printf("Password(Len: %v): %v \n", len(data.Password), data.Password)

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
	data, problems, err := jsonutils.DecodeValidJson[user.LoginUserReq](r)
	if err != nil {
		_ = jsonutils.Encode(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.AuthenticateUser(r.Context(), data.Email, data.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			jsonutils.Encode(w, r, http.StatusBadRequest, map[string]any{
				"error": "invalid email or password",
			})
			return
		}
		fmt.Println(err)
		jsonutils.Encode(w, r, http.StatusOK, map[string]any{
			"error": "unexpected internal server error",
		})
		return
	}
	//? middleware global que lida com todas as reqs
	err = api.Sessions.RenewToken(r.Context())
	fmt.Println("renew token error: ", err)
	if err != nil {
		jsonutils.Encode(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected server error",
		})
		return
	}

	api.Sessions.Put(r.Context(), "AuthenticatedUserId", id)

	jsonutils.Encode(w, r, http.StatusOK, map[string]any{
		"message": "user logged in successfully",
	})
}

func (api *API) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout route")
	err := api.Sessions.RenewToken(r.Context())
	fmt.Println("renew token error: ", err)
	if err != nil {
		jsonutils.Encode(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected server error",
		})
		return
	}

	api.Sessions.Remove(r.Context(), "AuthenticatedUserId")
	jsonutils.Encode(w, r, http.StatusOK, map[string]any{
		"message": "user logged out successfully",
	})
}
