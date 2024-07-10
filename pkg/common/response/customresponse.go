package response

type ExecutiveCommitteeResponse struct {
	ID          string `json:"id"`
	Role        string `json:"role"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Designation string `json:"designation"`
	Image       string `json:"image"`
}

type AlumniInfoForUser struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	StudentId      string `json:"student_id"`
	Email          string `json:"email"`
	GraduationYear string `json:"graduation_year"`
	Session        string `json:"session"`
	Role           string `json:"role"`
	ImagePath      string `json:"image_path"`
	JobType        string `json:"job_type"`
	SubJobType     string `json:"sub_job_type"`
	InstituteName  string `json:"institute_name"`
	JobTitle       string `json:"job_title"`
	PhoneNumber    string `json:"phone_number"`
	LinkedIn       string `json:"linked_in"`
	Facebook       string `json:"facebook"`
}
