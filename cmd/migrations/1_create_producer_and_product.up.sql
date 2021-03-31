-- download uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Producer and Product table defenitions
CREATE TABLE IF NOT EXISTS producers (
	id SERIAL UNIQUE NOT NULL,
	uuid uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
	producer_name TEXT NOT NULL,
	can_deliver BOOLEAN NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id, uuid)
);
CREATE TABLE IF NOT EXISTS products (
	id SERIAL UNIQUE NOT NULL,
	uuid uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4 (),
	product_name TEXT NOT NULL,
	-- Owner from producers table
	owner_id INTEGER NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now (),
	PRIMARY KEY (id, uuid),
	CONSTRAINT fk_producer_id FOREIGN KEY(owner_id) REFERENCES producers(id)
);