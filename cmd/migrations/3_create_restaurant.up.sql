-- Create Restaurant table

CREATE TABLE IF NOT EXISTS restaurants (
	id SERIAL UNIQUE NOT NULL,	
	uuid uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
	restaurant_name TEXT NOT NULL,
	PRIMARY KEY (id, uuid)
);

