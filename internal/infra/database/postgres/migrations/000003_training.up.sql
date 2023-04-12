CREATE TABLE IF NOT EXISTS trainings (
    id_training  UUID NOT NULL PRIMARY KEY,
	id_user UUID NOT NULL,
    trainig_name TEXT NOT NULL
)