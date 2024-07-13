package repositories

import (
	"alumni-management-server/pkg/models"
	"gorm.io/gorm"
)

type EventRepoInterface interface {
	Create(event *models.Event) (*models.Event, error)
	Update(event *models.Event) (*models.Event, error)
	EventCheck(title, startTime string) error
	FindById(id int) (models.Event, error)
}

type EventRepo struct {
	db *gorm.DB
}

func NewEventRepo(db *gorm.DB) EventRepo {
	return EventRepo{db: db}
}

func (eventRepo *EventRepo) Create(event *models.Event) (*models.Event, error) {
	if err := eventRepo.db.Create(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (eventRepo *EventRepo) Update(event *models.Event) (*models.Event, error) {
	if err := eventRepo.db.Table("events").Save(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (eventRepo *EventRepo) EventCheck(title, startTime string) error {
	query := eventRepo.db.Where("title = ? and start_time = ?", title, startTime).First(&models.Event{})
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (eventRepo *EventRepo) FindById(id int) (models.Event, error) {
	event := &models.Event{}
	if err := eventRepo.db.Table("events").Where("id = ?", id).First(&event).Error; err != nil {
		return *event, err
	}
	return *event, nil

}
