package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/reangeline/workout-plan-go/internal/infra/http/controllers"
)

func InitializeExerciseRoutes(uc *controllers.ExerciseController) chi.Router {
	r := chi.NewRouter()

	r.Route("/exercises", func(r chi.Router) {
		r.Post("/", uc.CreateExercise)
	})

	return r
}
