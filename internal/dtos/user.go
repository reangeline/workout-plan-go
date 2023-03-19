package dtos

import "github.com/reangeline/workout-plan-go/pkg/entities"

type CreateUserInput struct {
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
}

type UserOutputDTO struct {
	IDUser         entities.ID `json:"id"`
	FullName       string      `json:"full_name"`
	Email          string      `json:"email"`
	ProfilePicture string      `json:"profile_picture"`
}
