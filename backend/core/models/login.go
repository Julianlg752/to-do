package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

type LoginRequest struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    int64  `json:"userId"`
	Token string `json:"token"`
	Error string `json:"error"`
}

func (r *LoginRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.User, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}
