package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type application struct {
	firebase *firebase.App
}

func main() {
	fmt.Println(os.Getenv("FIREBASE_SERVICE_ACCOUNT"))
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_SERVICE_ACCOUNT"))
	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	app := application{
		firebase: firebaseApp,
	}

	createUser := flag.Bool("create", false, "flag to create a user")
	email := flag.String("email", "", "email of new user")
	password := flag.String("password", "", "password of new user")
	role := flag.String("role", "", "role of new user")
	flag.Parse()

	switch {
	case *createUser == true:

		if *email == "" || *password == "" {
			log.Fatal("password and email could not be blank")
		}

		err := app.createUser(*email, *password, role)
		if err != nil {
			log.Fatalf("could not create user %v\n", err)
		}
		log.Print("user created successfully")
	default:
		log.Print("No flags provided")
	}

}
