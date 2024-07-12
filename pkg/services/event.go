package services

import (
	"alumni-management-server/pkg/common/logger"
	"alumni-management-server/pkg/common/response"
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/repositories"
	"alumni-management-server/pkg/serializer"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type EventServiceInterface interface {
	Create(req *serializer.CreateEventRequest) error
}

type EventService struct {
	eventRepo repositories.EventRepoInterface
}

func NewEventService(eventRepo repositories.EventRepoInterface) EventService {
	return EventService{eventRepo: eventRepo}
}

func (eventService *EventService) Create(req *serializer.CreateEventRequest) error {
	if err := req.ValidateCreateEventRequest(); err != nil {
		logger.Error(err)
		return err
	}

	if err := eventService.eventRepo.EventCheck(req.Title, req.EventDate); err == nil {
		return response.ErrEventAlreadyExists
	}

	file, err := req.Image.Open()
	if err != nil {
		logger.Error(err)
		return err
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
	imagePath := filepath.Join(dirPath, strconv.FormatInt(time.Now().Unix(), 10)+"_"+req.Image.Filename)

	// Create the directory if it doesn't exist
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}

	dst, err := os.Create(imagePath)
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			return
		}
	}(dst)

	// Copy the uploaded file to the new file
	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	event := models.Event{}
	event.ToEventModel(req)

	event.ImagePath = imagePath

	if err := eventService.eventRepo.Create(&event); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
