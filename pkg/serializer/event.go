package serializer

import (
	v "github.com/go-ozzo/ozzo-validation"
	"mime/multipart"
)

type CreateEventRequest struct {
	Image       *multipart.FileHeader `json:"image"`
	Title       string                `json:"title"`
	EventDate   string                `json:"event_date"`
	StartTime   string                `json:"start_time"`
	Location    string                `json:"location"`
	Description string                `json:"description"`
}

func (r CreateEventRequest) ValidateCreateEventRequest() error {
	return v.ValidateStruct(&r,
		v.Field(&r.Title, v.Required.Error("Title cannot be empty")),
		v.Field(&r.EventDate, v.Required.Error("Event Date cannot be empty")),
		v.Field(&r.StartTime, v.Required.Error("Start Time cannot be empty")),
		v.Field(&r.Location, v.Required.Error("Location cannot be empty")),
		v.Field(&r.Image, v.Required.Error("Image cannot be empty")),
		v.Field(&r.Description, v.Required.Error("Description cannot be empty")),
	)
}
