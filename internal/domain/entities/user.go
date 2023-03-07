package entities

import "errors"

type User struct {
	ID             string
	FullName       string
	Email          string
	ProfilePicture string
}

func NewUser(full_name string, email string, profile_picture string) (*User, error) {
	user := &User{
		ID:             "1",
		FullName:       full_name,
		Email:          email,
		ProfilePicture: profile_picture,
	}

	err := user.IsValid()
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (u *User) IsValid() error {
	if u.ID == "" {
		return errors.New("id is required")
	}
	if u.FullName == "" {
		return errors.New("name is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	return nil
}
