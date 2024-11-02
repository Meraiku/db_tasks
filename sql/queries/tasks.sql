-- name: TaskThreeGetITEmployees :many
SELECT CONCAT(first_name, ' ', last_name) AS full_name
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

-- name: TaskSevenAggFunc :many
SELECT 
	departments.title AS department,
  COUNT(employees.id) AS total_employees
FROM employees
INNER JOIN departments ON employees.department_id = departments.id
GROUP BY departments.title
ORDER BY total DESC;
