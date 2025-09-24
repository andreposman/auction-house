package user

import (
	"context"

	"github.com/andreposman/action-house-api/internal/validator"
)

type CreateUserReq struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (req CreateUserReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.UserName), "user_name", "this field can't be empty")
	eval.CheckField(validator.MinChars(req.UserName, 5) &&
		validator.MaxChars(req.UserName, 50), "user_name", "this field should be between 5 and 50 chars")

	eval.CheckField(validator.MinChars(req.Email, 5) &&
		validator.MaxChars(req.Email, 100), "email", "this field should be between 5 and 100 chars")
	eval.CheckField(validator.Matches(req.Email, validator.EmailRegex), "email", "must be a valid email ")
	eval.CheckField(validator.NotBlank(req.Email), "email", "this field can't be empty")

	eval.CheckField(validator.NotBlank(req.Bio), "bio", "this field can't be empty")
	eval.CheckField(validator.MinChars(req.Bio, 5) &&
		validator.MaxChars(req.Bio, 255), "bio", "this field should be between 5 and 255 chars")

	return eval
}
