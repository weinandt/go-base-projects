package main

import (
	"fmt"

	"github.com/weinandt/go-base-projects/postgres/postgres"
	"github.com/weinandt/go-base-projects/postgres/user"
)

func setUpDependencies() (*user.UserInteractor, error) {
	db, err := postgres.CreatePostgresDatabase()
	if err != nil {
		return nil, err
	}

	userGateway := user.NewPostgresUserGateway(db)
	userInteractor := user.NewUserInteractor(userGateway)

	return userInteractor, nil
}

func main() {
	userInteractor, err := setUpDependencies()
	if err != nil {
		panic("Could not set up dependencies")
	}

	// Create a user in the db.
	newUser, err := userInteractor.CreateNewUser("nick")
	if err != nil {
		fmt.Println("failed to create user")
	} else {
		fmt.Println(newUser)
	}
}
