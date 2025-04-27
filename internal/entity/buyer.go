package entity

import (
	"github.com/google/uuid"
)

type Buyer struct {
	BuyerId     string `json:"buyer_id"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func NewBuyer(email, name, phoneNumber, password string) *Buyer {
	return &Buyer{
		BuyerId:     uuid.New().String(),
		Email:       email,
		FullName:    name,
		PhoneNumber: phoneNumber,
		Password:    password,
	}
}
