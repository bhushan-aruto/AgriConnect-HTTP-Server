package entity

import (
	"github.com/google/uuid"
)

type Farmer struct {
	FarmerId    string `json:"farmer_id"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func NewFarmer(email, name, phoneNumber, password string) *Farmer {
	return &Farmer{
		FarmerId:    uuid.New().String(),
		Email:       email,
		FullName:    name,
		PhoneNumber: phoneNumber,
		Password:    password,
	}
}
