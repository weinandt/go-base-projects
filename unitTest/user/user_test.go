package user_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/weinandt/go-base-projects/unitTest/user"
)

// Tests a business logic case.
// Does not actually call the mocked/stubbed gateway method.
func TestUserNameLength(t *testing.T) {
	// Arrange.
	// Using an embeded struct to mock.
	type MockUserGateway struct{ user.UserGateway }
	var mockUserGateway MockUserGateway
	userInteractor := user.NewUserInteractor(mockUserGateway)

	// Act.
	_, err := userInteractor.CreateNewUser("12")

	// Assert.
	if err == nil {
		t.Error("Expected an error to be returned.")
	}
}

// Notice the lower case vs Uppercase here. This allows an implementation to be swapped in the tests.
type StubbedUserGateway struct {
	createUser func(name string) (*user.User, error)
}

func (sug *StubbedUserGateway) CreateUser(name string) (*user.User, error) {
	return sug.createUser(name)
}

// Actually calls the gateway, so a stubbed implementation is necessary
func TestUserWithGatewayStubbed(t *testing.T) {
	// Arrange.
	expectedId := uuid.New().String()
	expectedUserName := "testName"
	myStubbedUserGateway := &StubbedUserGateway{
		createUser: func(name string) (*user.User, error) {
			return &user.User{
				Id:   expectedId,
				Name: expectedUserName,
			}, nil
		},
	}

	userInteractor := user.NewUserInteractor(myStubbedUserGateway)

	// Act
	newUser, err := userInteractor.CreateNewUser(expectedUserName)

	// Assert
	if err != nil {
		t.Fatal("Should not have returned error.")
	}

	isNameEqual := newUser.Name == expectedUserName
	isIdEqual := newUser.Id == expectedId

	if !isNameEqual || !isIdEqual {
		t.Fatal("Expected results do not match user.")
	}
}
