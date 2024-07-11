package services

import (
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/repositories"
)

type EventServiceInterface interface {
	Create(event *models.Event) error
}

type EventService struct {
	eventRepo repositories.EventRepoInterface
}

func NewEventService(eventRepo repositories.EventRepoInterface) EventService {
	return EventService{eventRepo: eventRepo}
}

func (eventService *EventService) Create(event *models.Event) error {
	return nil
}
