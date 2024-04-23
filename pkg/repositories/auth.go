package repositories

import (
	"alumni-management-server/pkg/customerror"
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/models"
	"errors"
	"gorm.io/gorm"
)

// userRepo defines the methods of the domain.IUserRepo interface.
type userRepo struct {
	db *gorm.DB
}

// UserDBInstance returns a new instance of the userRepo struct.
func UserDBInstance(d *gorm.DB) domain.IAuthRepo {
	return &userRepo{
		db: d,
	}
}

// DuplicateUserChecker returns a user model by the username.
func (repo *userRepo) DuplicateUserChecker(StudentId *string, Email *string) error {
	user := &models.UserDetail{}
	if err := repo.db.Where("student_id = ?", StudentId).First(user).Error; err == nil {
		return &customerror.StudentIDExistsError{ID: *StudentId}
	}
	if err := repo.db.Where("email = ?", Email).First(user).Error; err == nil {
		return &customerror.EmailExistsError{Email: *Email}
	}
	return nil
}

// CreateUser creates a new user with given user details.
func (repo *userRepo) CreateUser(user *models.UserDetail) error {
	if err := repo.db.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("username already exists")
		}
		return err
	}
	return nil
}

func (repo *userRepo) FindAuthorizedUserByEmailOrStudentId(value interface{}) (*models.UserDetail, error) {
	user := &models.UserDetail{}
	if err := repo.db.Where("student_id = ? OR email = ?", value, value).First(user).Error; err != nil {
		return nil, err
	}

	if user.IsUserVerified == false {
		return nil, &customerror.UserNotVerifiedError{}
	}

	return user, nil
}
