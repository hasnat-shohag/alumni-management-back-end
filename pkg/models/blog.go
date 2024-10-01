package models

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
)

type Blog struct {
	gorm.Model
	UserId        uint      `json:"user_id"`
	ImagePath     string    `json:"image_path"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Category      string    `json:"category"`
	Tags          string    `json:"tags"`
	Status        string    `json:"status"`
	Comments      []Comment `json:"comments" gorm:"foreignKey:BlogPostID"`
	CommentsCount int       `json:"comments_count"`
	Likes         []Like    `json:"likes" gorm:"foreignKey:BlogPostID"`
	LikesCount    int       `json:"likes_count"`
	Views         int       `json:"views"`
}

type Comment struct {
	gorm.Model
	UserID     uint   `json:"user_id"`
	BlogPostID uint   `json:"blog_post_id"`
	Content    string `json:"content"`
}

type Like struct {
	gorm.Model
	UserID     uint `json:"user_id"`
	BlogPostID uint `json:"blog_post_id"`
}

func (x *Blog) ToBlogModel(source interface{}) error {
	sourceValue := reflect.Indirect(reflect.ValueOf(source))
	destinationValue := reflect.Indirect(reflect.ValueOf(x))

	for i := 0; i < destinationValue.NumField(); i++ {
		destinationFieldName := destinationValue.Type().Field(i).Name
		sourceFieldValue := sourceValue.FieldByName(destinationFieldName)

		if sourceFieldValue.IsValid() && sourceFieldValue.CanSet() {
			destinationFieldValue := destinationValue.Field(i)
			if destinationFieldValue.Type() == sourceFieldValue.Type() {
				destinationFieldValue.Set(sourceFieldValue)
			} else {
				return fmt.Errorf("type mismatch for field %s", destinationFieldName)
			}
		}
	}
	return nil
}
