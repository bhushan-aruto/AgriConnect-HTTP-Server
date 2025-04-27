package models

type FarmerSignUpRequest struct {
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type FarmerLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
