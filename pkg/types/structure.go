package types

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"mime/multipart"
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

type ResetPasswordRequest struct {
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type CompleteProfileRequest struct {
	Image         *multipart.FileHeader `json:"image"`
	JobType       string                `json:"job_type"`
	InstituteName string                `json:"institute_name"`
	JobTitle      string                `json:"job_title"`
	PhoneNumber   string                `json:"phone_number"`
	LinkedIn      string                `json:"linked_in,omitempty"`
	Facebook      string                `json:"facebook,omitempty"`
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

func (request ResetPasswordRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.NewPassword,
			validation.Required.Error("New Password cannot be empty"),
			validation.Length(8, 128)),
		validation.Field(&request.ConfirmPassword,
			validation.Required.Error("Confirm Password cannot be empty"),
			validation.Length(8, 128),
			validation.By(func(value interface{}) error {
				if request.NewPassword != request.ConfirmPassword {
					return errors.New("passwords do not match")
				}
				return nil
			})),
	)
}

func (request CompleteProfileRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Image, validation.Required.Error("Image cannot be empty"), validation.By(func(value interface{}) error {
			if request.Image != nil && request.Image.Size > 1024*1024 {
				return errors.New("image size must be less than or equal to 1MB")
			}
			return nil
		})),
		validation.Field(&request.JobType, validation.Required.Error("Job Type cannot be empty")),
		validation.Field(&request.InstituteName, validation.Required.Error("Institute Name cannot be empty")),
		validation.Field(&request.JobTitle, validation.Required.Error("Job Title cannot be empty")),
		validation.Field(&request.PhoneNumber, validation.Required.Error("Phone Number cannot be empty")),
		validation.Field(&request.LinkedIn, validation.Length(0, 128)),
		validation.Field(&request.Facebook, validation.Length(0, 128)),
	)
}
