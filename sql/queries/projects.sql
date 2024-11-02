-- name: AddProject :exec
INSERT INTO projects(title) 
VALUES($1);
