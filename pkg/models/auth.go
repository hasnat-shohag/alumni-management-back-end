package models

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserDetail struct {
	gorm.Model
	StudentId      string    `gorm:"uniqueIndex;size:10"`
	Email          string    `gorm:"uniqueIndex;size:128"`
	PasswordHash   string    `json:"password_hash"`
	Name           string    `json:"name"`
	Session        string    `json:"session"`
	IsUserVerified bool      `json:"is_user_verified"`
	Role           string    `json:"role"`
	OTP            string    `json:"otp"`
	OtpExpiryTime  time.Time `json:"otp_expiry_time"`
	ImagePath      string    `json:"image_path"`
	JobType        string    `json:"job_type"`
	SubJobType     string    `json:"sub_job_type"`
	InstituteName  string    `json:"institute_name"`
	JobTitle       string    `json:"job_title"`
}

func (x *UserDetail) SetVerificationProperties() {
	x.IsUserVerified = false
	x.Session = generateSession(x.StudentId)
	x.Role = "user"
}

func generateSession(StudentId string) string {
	FirstNumber := StudentId[0:2]
	FirstNumberInt, err := strconv.Atoi(FirstNumber)
	if err != nil {
		return ""
	}
	SecondNumber := strconv.Itoa(FirstNumberInt - 1)
	return fmt.Sprintf("20%s-%s", SecondNumber, FirstNumber)
}
