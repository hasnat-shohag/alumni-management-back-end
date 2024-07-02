package models

import "gorm.io/gorm"

type ExecutiveCommittee struct {
	gorm.Model
	Role        string `json:"role"`
	Name        string `json:"name"`
	Designation string `json:"designation"`
}
