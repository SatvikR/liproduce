-- wipe old data (stale)
DELETE FROM products;
DELETE FROM producers;
DELETE FROM restaurants;

-- create user_id relations

ALTER TABLE producers
ADD COLUMN user_id INTEGER NOT NULL;

ALTER TABLE producers
ADD CONSTRAINT fk_producer_user_id FOREIGN KEY(user_id) REFERENCES users(id);

ALTER TABLE restaurants
ADD COLUMN user_id INTEGER NOT NULL;

ALTER TABLE restaurants
ADD CONSTRAINT fk_restaurant_user_id FOREIGN KEY(user_id) REFERENCES users(id);
