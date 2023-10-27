package main

import (
	"context"
	"fmt"

	"github.com/samverrall/errmaps/internal/core/services/users"
)

func main() {
	userSvc := users.NewService()

	_, err := userSvc.Signup(context.Background(), users.SignupReq{
		Email:    "test@test.com",
		Name:     "John Doe",
		Password: "supersecurepassword",
	})
	if err != nil {
		fmt.Println(err)
	}
}
