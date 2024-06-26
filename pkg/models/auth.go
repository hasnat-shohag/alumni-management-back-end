package models

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserDetail struct {
	gorm.Model
	Name                           string    `json:"name"`
	StudentId                      string    `gorm:"uniqueIndex;size:10"`
	Email                          string    `gorm:"uniqueIndex;size:128"`
	GraduationYear                 string    `json:"graduation_year"`
	PasswordHash                   string    `json:"password_hash"`
	Session                        string    `json:"session"`
	IsUserVerified                 bool      `json:"is_user_verified"`
	Role                           string    `json:"role"`
	OTP                            string    `json:"otp"`
	OtpExpiryTime                  time.Time `json:"otp_expiry_time"`
	ImagePath                      string    `json:"image_path"`
	CertificateOrStudentIdCardPath string    `json:"certificate_or_student_id_card_path"`
	JobType                        string    `json:"job_type"`
	SubJobType                     string    `json:"sub_job_type"`
	InstituteName                  string    `json:"institute_name"`
	JobTitle                       string    `json:"job_title"`
	PhoneNumber                    string    `json:"phone_number"`
	LinkedIn                       string    `json:"linked_in"`
	Facebook                       string    `json:"facebook"`
}

func (x *UserDetail) SetVerificationProperties() {
	x.IsUserVerified = false
	x.Session = generateSession(x.StudentId)
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
