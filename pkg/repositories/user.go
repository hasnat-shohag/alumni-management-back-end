package repositories

import (
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/utils"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
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
	// generate otp
	otp := rand.Int() % 1000000
	otpString := fmt.Sprintf("%06d", otp)

	// make otp hashed
	hashedOtp, err := utils.GetPasswordHash(otpString)
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

	return otpString, nil
}

func (repo *userRepo) UpdateUser(user *models.UserDetail) error {
	if err := repo.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *userRepo) FindAllAlumni(offset, limit int) ([]models.UserDetail, error) {
	var alumni []models.UserDetail
	if err := repo.db.Where("is_user_verified = ?", true).Offset(offset).Limit(limit).Find(&alumni).Error; err != nil {
		return nil, err
	}
	return alumni, nil
}

func (repo *userRepo) CountAuthorizedUser() (int, error) {
	var count int64
	if err := repo.db.Model(&models.UserDetail{}).Where("is_user_verified = ?", true).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
