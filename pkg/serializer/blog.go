package serializer

import (
	v "github.com/go-ozzo/ozzo-validation"
	"mime/multipart"
)

type CreateBlogRequest struct {
	UserId   uint                  `json:"user_id"`
	Image    *multipart.FileHeader `json:"image"`
	Title    string                `json:"title"`
	Content  string                `json:"content"`
	Category string                `json:"category"`
	Tags     string                `json:"tags"`
	Status   string                `json:"status"`
}

func (r CreateBlogRequest) ValidateCreateBlogRequest() error {
	return v.ValidateStruct(&r,
		v.Field(&r.Title, v.Required.Error("Title cannot be empty")),
		v.Field(&r.Content, v.Required.Error("Content cannot be empty")),
		//v.Field(&r.Category, v.Required.Error("Category cannot be empty")),
		//v.Field(&r.Tags, v.Required.Error("Tags cannot be empty")),
		//v.Field(&r.Status, v.Required.Error("Status cannot be empty")),
		v.Field(&r.Image, v.Required.Error("Image cannot be empty")),
	)
}
