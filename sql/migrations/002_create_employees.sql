-- +goose Up
CREATE TABLE IF NOT EXISTS employees(
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  department_id SERIAL NOT NULL REFERENCES departments(id)
);

-- +goose Down
DROP TABLE IF EXISTS employees;
