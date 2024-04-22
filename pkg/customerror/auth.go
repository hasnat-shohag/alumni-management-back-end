package customerror

import "fmt"

type StudentIDExistsError struct {
	ID string
}

func (e *StudentIDExistsError) Error() string {
	return fmt.Sprintf("Student ID %s already exists", e.ID)
}

type EmailExistsError struct {
	Email string
}

func (e *EmailExistsError) Error() string {
	fmt.Println(e.Email)
	return fmt.Sprintf("Email %s already exists", e.Email)
}

type UserNotVerifiedError struct{}

func (e *UserNotVerifiedError) Error() string {
	return "Please wait until verified"
}
