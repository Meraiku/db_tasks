-- name: TaskThreeGetITEmployees :many
SELECT (first_name, last_name) AS full_name
FROM employees
INNER JOIN departments ON employees.department_id = departments.id
WHERE departments.title = 'IT';

-- name: TastFourSetEmployee :exec
UPDATE employees
SET first_name = 'Robert'
WHERE id = 1;

-- name: TaskFiveDeleteProject :exec
DELETE FROM projects
WHERE id = 2;
