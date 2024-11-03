-- name: TaskThreeGetITEmployees :many
SELECT CONCAT(first_name, ' ', last_name) AS full_name
FROM employees e
INNER JOIN departments d ON e.department_id = d.id
WHERE d.title = 'IT';

-- name: TastFourSetEmployee :exec
UPDATE employees
SET first_name = 'Robert'
WHERE id = 1;

-- name: TaskFiveDeleteProject :exec
DELETE FROM projects
WHERE id = 2;

-- name: TaskSevenAggFunc :many
SELECT 
	d.title AS department,
  COUNT(e.id) AS total_employees
FROM employees e
INNER JOIN departments d ON e.department_id = d.id
GROUP BY d.title
ORDER BY total_employees DESC;

-- name: TaskEightJoin :many
SELECT
  e.first_name,
  e.last_name,
  d.title AS department
FROM employees e
INNER JOIN departments d ON e.department_id = d.id
ORDER BY d.title;
