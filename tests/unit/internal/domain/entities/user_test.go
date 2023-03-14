package entities_test

import (
	"testing"

	"github.com/reangeline/workout-plan-go/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyFields_WhenCreateUser_ThenShouldReceiveAnError(t *testing.T) {

	user := entities.User{
		FullName:       "",
		Email:          "",
		ProfilePicture: "",
	}
	err := user.IsValid()
	assert.Error(t, err)
}

func TestGivenAValidParams_WhenICallNewUser_ThenIShouldReceiveCreateUserWithAllParams(t *testing.T) {

	user := entities.User{
		FullName:       "Renato",
		Email:          "reangeline@hotmail.com",
		ProfilePicture: "http://image.jpge",
	}

	assert.Equal(t, "Renato", user.FullName)
	assert.Equal(t, "reangeline@hotmail.com", user.Email)
	assert.Equal(t, "http://image.jpge", user.ProfilePicture)

	assert.Nil(t, user.IsValid())

}

func TestGivenAValidParams_WhenICallNewUserFunc_ThenIShouldReceiveCreateUserWithAllParams(t *testing.T) {

	u, err := entities.NewUser("Renato", "reangeline@hotmail.com", "http://image.jpge")
	assert.Nil(t, err)

	assert.Equal(t, "Renato", u.FullName)
	assert.Equal(t, "reangeline@hotmail.com", u.Email)
	assert.Equal(t, "http://image.jpge", u.ProfilePicture)

}
