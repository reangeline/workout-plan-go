package entities

import (
	"errors"

	"github.com/reangeline/workout-plan-go/pkg/entities"
)

type User struct {
	IDUser         entities.ID
	FullName       string
	Email          string
	ProfilePicture string
}

func NewUser(full_name string, email string, profile_picture string) (*User, error) {
	user := &User{
		FullName:       full_name,
		Email:          email,
		ProfilePicture: profile_picture,
	}

	user.AddId()

	err := user.IsValid()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) AddId() {
	u.IDUser = entities.NewID()
}

func (u *User) IsValid() error {

	if u.FullName == "" {
		return errors.New("name is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	return nil
}
