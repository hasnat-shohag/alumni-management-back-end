package repositories

import (
	"alumni-management-server/pkg/models"
	"gorm.io/gorm"
)

type EventRepoInterface interface {
	Create(event *models.Event) error
	EventCheck(title, startTime string) error
}

type EventRepo struct {
	db *gorm.DB
}

func NewEventRepo(db *gorm.DB) EventRepo {
	return EventRepo{db: db}
}

func (eventRepo *EventRepo) Create(event *models.Event) error {
	if err := eventRepo.db.Create(event).Error; err != nil {
		return err
	}
	return nil
}

func (eventRepo *EventRepo) EventCheck(title, startTime string) error {
	query := eventRepo.db.Where("title = ? and start_time = ?", title, startTime).First(&models.Event{})
	if query.Error != nil {
		return query.Error
	}
	return nil
}
