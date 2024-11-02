-- +goose Up
CREATE TABLE IF NOT EXISTS projects(
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS projects;
