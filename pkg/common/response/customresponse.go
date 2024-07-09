package response

type ExecutiveCommitteeResponse struct {
	ID          string `json:"ID"`
	Role        string `json:"role"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Designation string `json:"designation"`
	Image       string `json:"image"`
}
