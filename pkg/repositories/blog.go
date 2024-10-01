package repositories

import (
	"alumni-management-server/pkg/models"
	"gorm.io/gorm"
)

type BlogRepoInterface interface {
	Create(blog *models.Blog) (*models.Blog, error)
}

type BlogRepo struct {
	db *gorm.DB
}

func NewBlogRepo(db *gorm.DB) BlogRepo {
	return BlogRepo{db: db}
}

func (blogRepo *BlogRepo) Create(blog *models.Blog) (*models.Blog, error) {
	if err := blogRepo.db.Create(blog).Error; err != nil {
		return nil, err
	}
	return blog, nil
}
