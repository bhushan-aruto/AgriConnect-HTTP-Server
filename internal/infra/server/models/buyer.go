package models

type BuyerSignUpRequest struct {
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type BuyerLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
