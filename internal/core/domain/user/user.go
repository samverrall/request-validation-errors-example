package user

import "errors"

type User struct{}

type Email string

func NewEmail(e string) (Email, error) {
	// Add validation here
	return "", errors.New("invalid email provided")
}

type Name string

func NewName(n string) (Name, error) {
	// Add validation here
	return "", errors.New("invalid name provided")
}

type Password string

func NewPassword(p string) (Password, error) {
	// Add validation here
	return "", errors.New("provided passport too short")
}
