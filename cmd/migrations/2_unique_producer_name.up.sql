-- Add unique constraint to producer table
ALTER TABLE producers
ADD CONSTRAINT unique_producer_name UNIQUE (producer_name);