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

func (repo *adminRepo) VerifyUser(studentId string, isValid bool) error {
	user := &models.UserDetail{}
	if err := repo.db.Where("student_id = ?", studentId).First(user).Error; err != nil {
		return err
	}

	if isValid == true {
		user.IsUserVerified = true
		if err := repo.db.Save(user).Error; err != nil {
			return err
		}
	} else {
		user.IsUserVerified = false
		if err := repo.db.Delete(user).Error; err != nil {
			return err
		}

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

func (repo *adminRepo) DeleteUser(user *models.UserDetail) error {

	if err := repo.db.Delete(user).Error; err != nil {
		return err
	}

	return nil
}

// AddExecutiveCommitteeMember adds a new executive committee member to the database.
func (repo *adminRepo) AddExecutiveCommitteeMember(executiveCommitteeMember *models.ExecutiveCommittee) error {
	if err := repo.db.Create(executiveCommitteeMember).Error; err != nil {
		return err
	}
	return nil
}

// DeleteExecutiveCommitteeMember deletes an executive committee member from the database.
func (repo *adminRepo) DeleteExecutiveCommitteeMember(execCommitteeMember *models.ExecutiveCommittee) error {
	if err := repo.db.Delete(execCommitteeMember).Error; err != nil {
		return err
	}
	return nil
}

// FindExecutiveCommitteeMemberById finds an executive committee member by ID.
func (repo *adminRepo) FindExecutiveCommitteeMemberById(id string) (models.ExecutiveCommittee, error) {
	execCommitteeMember := &models.ExecutiveCommittee{}
	if err := repo.db.Where("id = ?", id).First(execCommitteeMember).Error; err != nil {
		return *execCommitteeMember, err
	}
	return *execCommitteeMember, nil
}

// UpdateExecutiveCommitteeMember updates an executive committee member in the database.
func (repo *adminRepo) UpdateExecutiveCommitteeMember(execCommitteeMember *models.ExecutiveCommittee) error {
	if err := repo.db.Save(execCommitteeMember).Error; err != nil {
		return err
	}
	return nil
}

// GetExecutiveCommitteeInfo returns the executive committee information.
func (repo *adminRepo) GetExecutiveCommitteeInfo() ([]models.ExecutiveCommittee, error) {
	var executiveCommittee []models.ExecutiveCommittee
	if err := repo.db.Find(&executiveCommittee).Error; err != nil {
		return nil, err
	}

	return executiveCommittee, nil
}
