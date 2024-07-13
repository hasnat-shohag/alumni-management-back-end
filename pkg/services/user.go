package services

import (
	"alumni-management-server/pkg/common/response"
	"alumni-management-server/pkg/domain"
	"alumni-management-server/pkg/email"
	"alumni-management-server/pkg/models"
	"alumni-management-server/pkg/serializer"
	"alumni-management-server/pkg/utils"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

type userService struct {
	userRepo  domain.IUserRepo
	authRepo  domain.IAuthRepo
	adminRepo domain.IAdminRepo
}

func UserServiceInstance(userRepo domain.IUserRepo, authRepo domain.IAuthRepo, adminRepo domain.IAdminRepo) domain.IUserService {
	return &userService{
		userRepo:  userRepo,
		authRepo:  authRepo,
		adminRepo: adminRepo,
	}
}

func (userService *userService) ForgetPassword(Email string) error {
	// Check if user exists
	user, err := userService.authRepo.FindAuthorizedUserByEmailOrStudentId(Email)
	if err != nil {
		return fmt.Errorf("user with email %s does not exist", Email)
	}

	// Create OTP
	otp, err := userService.userRepo.CreateOTP(user)
	if err != nil {
		return err
	}

	// Send OTP to user
	link := fmt.Sprintf("http://localhost:9030/reset-password?otp=%s&email=%s", otp, Email)
	emailBody, err := email.CreateForgotPasswordEmail(link)
	if err != nil {
		return err
	}

	err = email.SendEmail(user.Email, email.PasswordResetSubject, emailBody)
	if err != nil {
		return err
	}
	return nil
}

func (userService *userService) GetUserFromEmailWithValidOtp(email, otp string) (*models.UserDetail, error) {
	user, err := userService.authRepo.FindAuthorizedUserByEmailOrStudentId(email)
	if err != nil {
		return nil, err
	}
	// check user otp is valid or not
	if err := utils.CheckPassword(user.OTP, otp); err != nil {
		return nil, fmt.Errorf("invalid otp")
	}

	// check otp expiry time
	if user.OtpExpiryTime.Before(user.OtpExpiryTime) {
		return nil, fmt.Errorf("otp expired, try again")
	}

	return user, nil
}

func (userService *userService) ResetPassword(user *models.UserDetail, password string) error {
	// Hash the password
	hashedPassword, err := utils.GetPasswordHash(password)
	if err != nil {
		return err
	}

	// Update the user password
	user.PasswordHash = hashedPassword
	err = userService.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (userService *userService) GetAllAlumni(page, limit int, jobType, instituteName string) ([]response.AlumniInfoForUser, int, error) {
	offset := (page - 1) * limit
	var alumni []models.UserDetail
	alumni, totalRecords, err := userService.userRepo.FindAllAlumni(offset, limit, jobType, instituteName)

	if err != nil {
		return nil, 0, err
	}

	var alumniResponse []response.AlumniInfoForUser
	for _, user := range alumni {
		var alumniInfo response.AlumniInfoForUser
		alumniInfo.ID = strconv.Itoa(int(user.ID))
		alumniInfo.Name = user.Name
		alumniInfo.StudentId = user.StudentId
		alumniInfo.Email = user.Email
		alumniInfo.GraduationYear = user.GraduationYear
		alumniInfo.Session = user.Session
		alumniInfo.Role = user.Role

		if user.ImagePath != "" {
			alumniInfo.ImagePath = utils.GetImageUrl(user.ImagePath)
		} else {
			alumniInfo.ImagePath = user.ImagePath
		}

		alumniInfo.JobType = user.JobType
		alumniInfo.InstituteName = user.InstituteName
		alumniInfo.JobTitle = user.JobTitle
		alumniInfo.PhoneNumber = user.PhoneNumber
		alumniInfo.LinkedIn = user.LinkedIn
		alumniInfo.Facebook = user.Facebook

		alumniResponse = append(alumniResponse, alumniInfo)
	}

	return alumniResponse, totalRecords, nil
}

func (userService *userService) GetAlumni(id string) (*response.AlumniInfoForUser, error) {
	user, err := userService.userRepo.FindAlumni(id)
	if err != nil {
		return nil, err
	}

	// customized response as unnecessary information not needed to client
	var customizedResponse response.AlumniInfoForUser

	customizedResponse.ID = strconv.Itoa(int(user.ID))
	customizedResponse.Name = user.Name
	customizedResponse.StudentId = user.StudentId
	customizedResponse.Email = user.Email
	customizedResponse.GraduationYear = user.GraduationYear
	customizedResponse.Session = user.Session
	customizedResponse.Role = user.Role

	if user.ImagePath != "" {
		customizedResponse.ImagePath = utils.GetImageUrl(user.ImagePath)
	} else {
		customizedResponse.ImagePath = user.ImagePath
	}

	customizedResponse.JobType = user.JobType
	customizedResponse.InstituteName = user.InstituteName
	customizedResponse.JobTitle = user.JobTitle
	customizedResponse.PhoneNumber = user.PhoneNumber
	customizedResponse.LinkedIn = user.LinkedIn
	customizedResponse.Facebook = user.Facebook

	return &customizedResponse, nil
}

func (userService *userService) DeleteMe(studentId, studentIdFromToken string) error {
	if studentId != studentIdFromToken {
		return fmt.Errorf("you have no access to delete others account")
	}

	// Find the alumni or student, coz student can also delete their account
	user, err := userService.userRepo.FindUser(studentId)
	if err != nil {
		return nil
	}

	err = userService.adminRepo.DeleteUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (userService *userService) UpdateMe(studentId string, request serializer.CompleteProfileRequest) error {
	// check user exists
	user, err := userService.userRepo.FindAlumni(studentId)
	if err != nil {
		return err
	}

	// Open the alumni image file
	file, err := request.Image.Open()
	if err != nil {
		return err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// Create a new file in the desired location
	dirPath := "./images/alumni_avatar/"
	imagePath := filepath.Join(dirPath, studentId+"_"+request.Image.Filename)

	// Create the directory if it doesn't exist
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}

	dst, err := os.Create(imagePath)
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			return
		}
	}(dst)

	// Copy the uploaded file to the new file
	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	// Update the user's image with the image from the request
	user.ImagePath = imagePath
	user.JobType = request.JobType
	user.InstituteName = request.InstituteName
	user.JobTitle = request.JobTitle
	user.PhoneNumber = request.PhoneNumber
	if request.LinkedIn != "" {
		user.LinkedIn = request.LinkedIn
	}
	if request.Facebook != "" {
		user.Facebook = request.Facebook

	}

	// Save the updated user back to the database
	err = userService.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}
