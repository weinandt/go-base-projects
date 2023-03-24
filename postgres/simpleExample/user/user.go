package user

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
	return userInteractor.userGateway.createUser(name)
}
