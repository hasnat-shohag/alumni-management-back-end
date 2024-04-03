package models

import (
	"gorm.io/gorm"
)

type UserDetail struct {
	gorm.Model
	StudentId      string `gorm:"uniqueIndex;size:10"`
	Email          string `gorm:"uniqueIndex;size:128"`
	PasswordHash   string `json:"password_hash"`
	Name           string `json:"name"`
	IsUserVerified bool   `json:"is_user_verified"`
	Role           string `json:"role"`
}

func (x *UserDetail) SetVerificationProperties() {
	x.IsUserVerified = false
	x.Role = "user"
}
