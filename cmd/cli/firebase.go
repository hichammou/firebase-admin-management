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

	userRecord, err := authClient.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}

	if *role != "" {
		claims := map[string]interface{}{
			"role": role,
		}
		err = authClient.SetCustomUserClaims(context.Background(), userRecord.UID, claims)
		if err != nil {
			log.Println("could not set user role")
		}
	}

	return nil
}
