CREATE TABLE IF NOT EXISTS exercises (
    id_exercise  UUID NOT NULL PRIMARY KEY,
	name_exercise TEXT NOT NULL,
	description TEXT NOT NULL,
    uri_gif TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS muscle_group_exercises (
    id_exercise  UUID NOT NULL,
	muscle_group TEXT NOT NULL,
    CONSTRAINT fk_exercise
      FOREIGN KEY(id_exercise) 
	  REFERENCES exercises(id_exercise)
)

