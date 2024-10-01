package services

import (
	"alumni-management-server/pkg/common/logger"
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/repositories"
	"alumni-management-server/pkg/serializer"
	"alumni-management-server/pkg/utils"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

type BlogServiceInterface interface {
	Create(request *serializer.CreateBlogRequest) (*models.Blog, error)
}

type BlogService struct {
	blogRepo repositories.BlogRepoInterface
}

func NewBlogService(blogRepo repositories.BlogRepoInterface) BlogService {
	return BlogService{blogRepo: blogRepo}
}

func (blogService *BlogService) Create(request *serializer.CreateBlogRequest) (*models.Blog, error) {
	if err := request.ValidateCreateBlogRequest(); err != nil {
		logger.Error(err)
		return nil, err
	}

	file, err := request.Image.Open()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			logger.Error(err)
			return
		}
	}(file)

	// Create a new file in the desired location
	dirPath := "./images/blog_image"
	imagePath := filepath.Join(dirPath, strconv.FormatInt(utils.GenerateRandomNumberOfSixDigit(), 6)+"_"+request.Image.Filename)

	// Create the directory if it doesn't exist
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
	}

	dst, err := os.Create(imagePath)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			logger.Error(err)
			return
		}
	}(dst)

	// Copy the uploaded file to the new file
	if _, err := io.Copy(dst, file); err != nil {
		logger.Error(err)
		return nil, err
	}

	blog := models.Blog{}
	if err := blog.ToBlogModel(request); err != nil {
		logger.Error(err)
		return nil, err
	}

	fmt.Println("service calling OKK")
	blog.ImagePath = imagePath

	newBlog, err := blogService.blogRepo.Create(&blog)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return newBlog, nil
}
