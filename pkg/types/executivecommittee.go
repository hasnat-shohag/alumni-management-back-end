package types

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"mime/multipart"
)

type CreateCommitteeRequest struct {
	Role        string                `json:"role"`
	Name        string                `json:"name"`
	Email       string                `json:"email"`
	Designation string                `json:"designation"`
	Image       *multipart.FileHeader `json:"image"`
}

func (request CreateCommitteeRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Role, validation.Required.Error("Role cannot be empty")),
		validation.Field(&request.Name, validation.Required.Error("Name cannot be empty")),
		validation.Field(&request.Designation, validation.Required.Error("Designation cannot be empty")),
		validation.Field(&request.Email, validation.Required.Error("Email cannot be empty")),
		validation.Field(&request.Image, validation.Required.Error("Image cannot be empty")),
	)
}

type UpdateCommitteeRequest struct {
	CreateCommitteeRequest
}
