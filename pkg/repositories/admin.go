package repositories

import (
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/models"
	"gorm.io/gorm"
)

type adminRepo struct {
	db *gorm.DB
}

// AdminDBInstance returns a new instance of the adminRepo struct.
func AdminDBInstance(d *gorm.DB) domain.IAdminRepo {
	return &adminRepo{
		db: d,
	}
}

func (repo *adminRepo) VerifyUser(studentId string) error {
	user := &models.UserDetail{}
	if err := repo.db.Where("student_id = ?", studentId).First(user).Error; err != nil {
		return err
	}

	user.IsUserVerified = true
	if err := repo.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *adminRepo) FindUserByStudentId(studentId string) (models.UserDetail, error) {
	user := &models.UserDetail{}
	if err := repo.db.Where("student_id = ?", studentId).First(user).Error; err != nil {
		return *user, err
	}
	return *user, nil
}
