package main

import (
	"context"
	"log"

	"firebase.google.com/go/v4/auth"
)

func (app application) createUser(email, password string, role *string) error {
	authClient, err := app.firebase.Auth(context.Background())
	if err != nil {
		return err
	}

	user := (&auth.UserToCreate{}).Email(email).Password(password)
	if *role != "" {
		log.Println("role not empty")
	}
	_, err = authClient.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}
