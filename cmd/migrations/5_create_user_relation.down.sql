-- Drop constraints

ALTER TABLE producers
DROP CONSTRAINT fk_producer_user_id;

ALTER TABLE producers
DROP COLUMN user_id;

ALTER TABLE restaurants
DROP CONSTRAINT fk_restaurant_user_id;

ALTER TABLE restaurants
DROP COLUMN user_id;
