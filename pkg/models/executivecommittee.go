package models

import "gorm.io/gorm"

type ExecutiveCommittee struct {
	gorm.Model
	Role        string `json:"role"`
	Name        string `json:"name"`
	Email       string `gorm:"uniqueIndex,size:128"`
	Designation string `json:"designation"`
	ImagePath   string `json:"image_path"`
}
