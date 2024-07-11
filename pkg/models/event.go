package models

type Event struct {
	ID          int    `json:"id"`
	ImagePath   string `json:"image_path"`
	Title       string `json:"title"`
	EventDate   string `json:"event_date"`
	StartTime   string `json:"start_time"`
	Location    string `json:"location"`
	Description string `json:"description"`
	IsDeleted   bool   `json:"is_deleted" gorm:"default:false"`
}
