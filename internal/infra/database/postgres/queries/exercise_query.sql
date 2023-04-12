-- name: CreateExercise :exec
INSERT INTO exercises (id_exercise, name_exercise, description, uri_gif) 
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: CreateMuscleGroupExercises :exec
INSERT INTO muscle_group_exercises (id_exercise, muscle_group) 
VALUES ($1,$2)
RETURNING *;