package database

import (
	"database/sql"
	"fmt"

	"github.com/reangeline/workout-plan-go/internal/domain/entities"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) CreateUser(user *entities.User) error {
	fmt.Println(user)
	return nil
}
