package repositories

import (
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/utils"
	"crypto/rand"
	"encoding/base64"
	"gorm.io/gorm"
	"time"
)

type userRepo struct {
	db *gorm.DB
}

// UserDBInstance returns a new instance of the userRepo struct.
func UserDBInstance(d *gorm.DB) domain.IUserRepo {
	return &userRepo{
		db: d,
	}
}

func (repo *userRepo) CreateOTP(user *models.UserDetail) (string, error) {
	otpBytes := make([]byte, 16)
	if _, err := rand.Read(otpBytes); err != nil {
		return "", err
	}

	otp := base64.URLEncoding.EncodeToString(otpBytes)
	// make otp hashed
	hashedOtp, err := utils.GetPasswordHash(otp)
	if err != nil {
		return "", err
	}
	expiryTime := time.Now().Add(5 * time.Minute)

	user.OTP = hashedOtp
	user.OtpExpiryTime = expiryTime

	if err := repo.UpdateUser(user); err != nil {
		return "", err
	}

	if user.OtpExpiryTime.IsZero() {
		user.OtpExpiryTime = time.Now()
	}

	return otp, nil
}

func (repo *userRepo) UpdateUser(user *models.UserDetail) error {
	if err := repo.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
