-- name: AddDepartment :one
INSERT INTO departments(title) 
VALUES($1)
RETURNING id;
