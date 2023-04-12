package database_test

import (
	"context"

	"github.com/reangeline/workout-plan-go/internal/domain/entities"
)

type MemoryExerciseRepository struct {
	Exercises map[string]*entities.Exercise
}

func NewMemoryExerciseRepository() *MemoryExerciseRepository {
	return &MemoryExerciseRepository{
		Exercises: make(map[string]*entities.Exercise),
	}
}

func (r *MemoryExerciseRepository) CreateExercise(ctx context.Context, exercise *entities.Exercise) error {
	r.Exercises[exercise.NameExercise] = exercise
	return nil
}
