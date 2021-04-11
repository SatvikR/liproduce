-- Create users table

CREATE TABLE IF NOT EXISTS users (
	id SERIAL UNIQUE NOT NULL,
	username TEXT UNIQUE NOT NULL ,
	user_type TEXT NOT NULL,
	email TEXT UNIQUE NOT NULL,
	"password" TEXT NOT NULL,
	PRIMARY KEY (id)
);
