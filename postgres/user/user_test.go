package user_test

import (
	"testing"

	"github.com/weinandt/go-base-projects/postgres/user"
)

func TestUserNameLength(t *testing.T) {
	// Arrange.
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
