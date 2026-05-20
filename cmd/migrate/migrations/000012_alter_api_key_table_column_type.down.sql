ALTER TABLE api_key 
ALTER COLUMN type TYPE VARCHAR(50) 
USING array_to_string(type, ',');