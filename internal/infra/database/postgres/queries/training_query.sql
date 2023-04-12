-- name: CreateTraining :exec
INSERT INTO trainings (id_training, id_user, trainig_name) 
VALUES ($1,$2,$3)
RETURNING *;
