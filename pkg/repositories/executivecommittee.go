package repositories

import (
	"alumni-management-server/pkg/models"
	"fmt"

	"gorm.io/gorm"
)

type ExecutiveCommitteeRepoInterface interface {
	AddExecutiveCommitteeMember(executiveCommitteeMember *models.ExecutiveCommittee) error
	DeleteExecutiveCommitteeMember(execCommitteeMember *models.ExecutiveCommittee) error
	UpdateExecutiveCommitteeMember(execCommitteeMember *models.ExecutiveCommittee) error
	GetExecutiveCommitteeInfo() ([]models.ExecutiveCommittee, error)
	FindBy(fieldName string, value any) (models.ExecutiveCommittee, error)
}

type ExecutiveCommitteeRepo struct {
	db *gorm.DB
}

func NewExecutiveCommitteeRepo(db *gorm.DB) ExecutiveCommitteeRepo {
	return ExecutiveCommitteeRepo{db: db}
}

// AddExecutiveCommitteeMember adds a new executive committee member to the database.
func (repo *ExecutiveCommitteeRepo) AddExecutiveCommitteeMember(executiveCommitteeMember *models.ExecutiveCommittee) error {
	if err := repo.db.Create(executiveCommitteeMember).Error; err != nil {
		return err
	}
	return nil
}

// DeleteExecutiveCommitteeMember deletes an executive committee member from the database.
func (repo *ExecutiveCommitteeRepo) DeleteExecutiveCommitteeMember(execCommitteeMember *models.ExecutiveCommittee) error {
	if err := repo.db.Delete(execCommitteeMember).Error; err != nil {
		return err
	}
	return nil
}

// FindExecutiveCommitteeMemberById finds an executive committee member by ID.
func (repo *ExecutiveCommitteeRepo) FindExecutiveCommitteeMemberById(id string) (models.ExecutiveCommittee, error) {
	execCommitteeMember := &models.ExecutiveCommittee{}
	if err := repo.db.Where("id = ?", id).First(execCommitteeMember).Error; err != nil {
		return *execCommitteeMember, err
	}
	return *execCommitteeMember, nil
}

// UpdateExecutiveCommitteeMember updates an executive committee member in the database.
func (repo *ExecutiveCommitteeRepo) UpdateExecutiveCommitteeMember(execCommitteeMember *models.ExecutiveCommittee) error {
	if err := repo.db.Save(execCommitteeMember).Error; err != nil {
		return err
	}
	return nil
}

// GetExecutiveCommitteeInfo returns the executive committee information.
func (repo *ExecutiveCommitteeRepo) GetExecutiveCommitteeInfo() ([]models.ExecutiveCommittee, error) {
	var executiveCommittee []models.ExecutiveCommittee
	if err := repo.db.Find(&executiveCommittee).Error; err != nil {
		return nil, err
	}

	return executiveCommittee, nil
}

// FindBy check user already exist or not
func (repo *ExecutiveCommitteeRepo) FindBy(fieldName string, value any) (models.ExecutiveCommittee, error) {
	executiveMember := &models.ExecutiveCommittee{}
	query := fmt.Sprintf("%s = ?", fieldName) // Dynamically create the query string
	if err := repo.db.Where(query, value).First(executiveMember).Error; err != nil {
		return *executiveMember, err
	}
	return *executiveMember, nil
}
