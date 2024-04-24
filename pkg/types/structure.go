package types

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

// SignupRequest defines the request body for the signup request.
type SignupRequest struct {
	StudentId string `json:"student_id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Password  string `json:"password"`
}

type LoginRequest struct {
	StudentId *string `json:"student_id,omitempty"`
	Email     *string `json:"email,omitempty"`
	Password  string  `json:"password"`
}

type LoginResponse struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	StudentId      string `json:"student_id"`
	IsUserVerified bool   `json:"is_user_verified"`
	IsActive       bool   `json:"is_active"`
	Role           string `json:"role"`
	AccessToken    string `json:"access_token"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

// Validate validates the request body for the SignupRequest.
func (request SignupRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.StudentId,
			validation.Required.Error("Student Id cannot be empty"),
			validation.Length(10, 10)),
		validation.Field(&request.Password,
			validation.Required.Error("Password cannot be empty"),
			validation.Length(8, 128)),
		validation.Field(&request.Name,
			validation.Required.Error("Name cannot be empty"),
			validation.Length(2, 64)),
		validation.Field(&request.Email,
			validation.Required.Error("Email cannot be empty"),
			validation.Length(4, 128)))
}

func (request LoginRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Password, validation.Required.Error("Password cannot be empty"), validation.Length(8, 128)),
		validation.Field(&request.StudentId, validation.By(func(value interface{}) error {
			if request.StudentId == nil && request.Email == nil {
				return errors.New("Must give StudentId or Email")
			}
			return nil
		})),
	)
}

func (request ForgotPasswordRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Email,
			validation.Required.Error("Email cannot be empty"),
			validation.Length(4, 128)))
}
