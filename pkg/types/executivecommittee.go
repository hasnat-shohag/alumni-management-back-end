package types

import validation "github.com/go-ozzo/ozzo-validation"

type CreateCommitteeRequest struct {
	Role        string `json:"role"`
	Name        string `json:"name"`
	Designation string `json:"designation"`
}

func (request CreateCommitteeRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Role, validation.Required.Error("Role cannot be empty")),
		validation.Field(&request.Name, validation.Required.Error("Name cannot be empty")),
		validation.Field(&request.Designation, validation.Required.Error("Designation cannot be empty")),
	)
}

type UpdateCommitteeRequest struct {
	CreateCommitteeRequest
}
