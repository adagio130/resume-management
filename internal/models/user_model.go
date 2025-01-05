package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id"`
	Account  string `json:"account"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Location string `json:"location"`
}

func NewUser(name, account, gender, location string) *User {
	userId := uuid.New().String()
	return &User{
		ID:       userId,
		Name:     name,
		Account:  account,
		Gender:   gender,
		Location: location,
	}
}
