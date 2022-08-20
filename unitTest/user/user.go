package user

import "errors"

type User struct {
	id   string
	name string
}

// This is the object which will contain unit testable business logic.
type UserInteractor struct {
	userGateway UserGateway
}

func NewUserInteractor(userGateway UserGateway) *UserInteractor {
	return &UserInteractor{
		userGateway: userGateway,
	}
}

func (userInteractor *UserInteractor) CreateNewUser(name string) (*User, error) {
	// Here is some fake business logic to prove the point on how this can be unit tested.
	// We won't allow names for users of fewer than 3 characters.
	if len(name) < 3 {
		return nil, errors.New("Name of a user must be at least 3 characters.")
	}

	return userInteractor.userGateway.createUser(name)
}
