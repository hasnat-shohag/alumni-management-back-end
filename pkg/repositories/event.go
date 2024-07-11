package repositories

import (
	"alumni-management-server/pkg/models"
	"gorm.io/gorm"
)

type EventRepoInterface interface {
	Create(event *models.Event) error
}

type EventRepo struct {
	db *gorm.DB
}

func NewEventRepo(db *gorm.DB) EventRepo {
	return EventRepo{db: db}
}

func (eventRepo *EventRepo) Create(event *models.Event) error {
	return nil
}
