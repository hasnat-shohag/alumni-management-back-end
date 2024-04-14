package domain

type IAdminRepo interface {
	VerifyUser(studentId string) error
}

type IAdminService interface {
	VerifyUser(studentId string) error
}
