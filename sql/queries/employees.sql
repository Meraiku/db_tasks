-- name: AddEmployee :one
INSERT INTO employees(first_name, last_name, department_id) 
VALUES($1, $2, $3)
RETURNING id;
