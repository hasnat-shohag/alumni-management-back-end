package repositories

import (
	"alumni-management-server/pkg/common/response"
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/models"
	"errors"
	"gorm.io/gorm"
)

// userRepo defines the methods of the domain.IUserRepo interface.
type authRepo struct {
	db *gorm.DB
}

// UserDBInstance returns a new instance of the userRepo struct.
func AuthDBInstance(d *gorm.DB) domain.IAuthRepo {
	return &authRepo{
		db: d,
	}
}

// DuplicateUserChecker returns a user model by the username.
func (repo *authRepo) DuplicateUserChecker(StudentId *string, Email *string) error {
	user := &models.UserDetail{}
	if err := repo.db.Where("student_id = ?", StudentId).First(user).Error; err == nil {
		return &response.StudentIDExistsError{ID: *StudentId}
	}
	if err := repo.db.Where("email = ?", Email).First(user).Error; err == nil {
		return &response.EmailExistsError{Email: *Email}
	}
	return nil
}

// CreateUser creates a new user with given user details.
func (repo *authRepo) CreateUser(user *models.UserDetail) error {
	if err := repo.db.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("username already exists")
		}
		return err
	}
	return nil
}

func (repo *authRepo) FindAuthorizedUserByEmailOrStudentId(value interface{}) (*models.UserDetail, error) {
	user := &models.UserDetail{}
	if err := repo.db.Where("student_id = ? OR email = ?", value, value).First(user).Error; err != nil {
		return nil, err
	}

	if user.IsUserVerified == false {
		return nil, &response.UserNotVerifiedError{}
	}

	return user, nil
}
