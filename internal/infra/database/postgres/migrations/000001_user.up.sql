CREATE TABLE IF NOT EXISTS users (
    id_user  UUID NOT NULL PRIMARY KEY,
	full_name TEXT NOT NULL,
	email TEXT NOT NULL,
	profile_picture TEXT
);