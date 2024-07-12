package models

import "reflect"

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

func (x *Event) ToEventModel(source interface{}) {
	sourceValue := reflect.Indirect(reflect.ValueOf(source))
	destinationValue := reflect.Indirect(reflect.ValueOf(x))

	for i := 0; i < destinationValue.NumField(); i++ {
		destinationFieldName := destinationValue.Type().Field(i).Name
		sourceFieldValue := sourceValue.FieldByName(destinationFieldName)

		if sourceFieldValue.IsValid() && sourceFieldValue.CanSet() {
			destinationFieldValue := destinationValue.Field(i)
			destinationFieldValue.Set(sourceFieldValue)
		}
	}

}
