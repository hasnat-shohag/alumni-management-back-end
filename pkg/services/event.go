package services

import (
	"alumni-management-server/pkg/common/logger"
	"alumni-management-server/pkg/common/response"
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/repositories"
	"alumni-management-server/pkg/serializer"
	"alumni-management-server/pkg/utils"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

type EventServiceInterface interface {
	Create(req *serializer.CreateEventRequest) (*models.Event, error)
	Update(id int, req *serializer.CreateEventRequest) (*models.Event, error)
	Delete(id int) (int, error)
	FindById(id int) (*models.Event, error)
	FindAll() (*[]models.Event, error)
}

type EventService struct {
	eventRepo repositories.EventRepoInterface
}

func NewEventService(eventRepo repositories.EventRepoInterface) EventService {
	return EventService{eventRepo: eventRepo}
}

func (eventService *EventService) Create(req *serializer.CreateEventRequest) (*models.Event, error) {
	if err := req.ValidateCreateEventRequest(); err != nil {
		logger.Error(err)
		return nil, err
	}

	if err := eventService.eventRepo.EventCheck(req.Title, req.EventDate); err == nil {
		return nil, response.ErrEventAlreadyExists
	}

	file, err := req.Image.Open()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			logger.Error(err)
			return
		}
	}(file)

	// Create a new file in the desired location
	dirPath := "./images/event_banner"
	imagePath := filepath.Join(dirPath, strconv.FormatInt(utils.GenerateRandomNumberOfSixDigit(), 6)+"_"+req.Image.Filename)

	// Create the directory if it doesn't exist
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return nil, err
		}
	}

	dst, err := os.Create(imagePath)
	if err != nil {
		return nil, err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			return
		}
	}(dst)

	// Copy the uploaded file to the new file
	if _, err := io.Copy(dst, file); err != nil {
		return nil, err
	}

	event := models.Event{}
	event.ToEventModel(req)

	event.ImagePath = imagePath

	newEvent, err := eventService.eventRepo.Create(&event)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return newEvent, nil
}

func (eventService *EventService) Update(id int, req *serializer.CreateEventRequest) (*models.Event, error) {
	if err := req.ValidateCreateEventRequest(); err != nil {
		logger.Error(err)
		return nil, err
	}

	event, err := eventService.eventRepo.FindById(id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	file, err := req.Image.Open()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			logger.Error(err)
			return
		}
	}(file)

	// Create a new file in the desired location
	dirPath := "./images/event_banner"
	imagePath := filepath.Join(dirPath, strconv.FormatInt(utils.GenerateRandomNumberOfSixDigit(), 6)+"_"+req.Image.Filename)

	dst, err := os.Create(imagePath)
	if err != nil {
		return nil, err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			return
		}
	}(dst)

	// Copy the uploaded file to the new file
	if _, err := io.Copy(dst, file); err != nil {
		return nil, err
	}

	event.ToEventModel(req)

	event.ID = id
	event.ImagePath = imagePath

	updatedEvent, err := eventService.eventRepo.Update(&event)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return updatedEvent, nil
}

func (eventService *EventService) Delete(id int) (int, error) {
	_, err := eventService.eventRepo.FindById(id)
	if err != nil {
		logger.Error(err)
		return 0, fmt.Errorf("error: %s", err)
	}

	ID, err := eventService.eventRepo.Delete(id)
	if err != nil {
		logger.Error(err)
		return 0, fmt.Errorf("error: %s", err)
	}

	return ID, nil
}

func (eventService *EventService) FindById(id int) (*models.Event, error) {
	event, err := eventService.eventRepo.FindById(id)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("error: %s", err)
	}
	return &event, nil
}

func (eventService *EventService) FindAll() (*[]models.Event, error) {
	events, err := eventService.eventRepo.FindAll()
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("error: %s", err)
	}
	return events, nil
}
