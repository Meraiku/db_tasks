package repo

import "github.com/meraiku/db_tasks/internal/models"

func FromModelToRepoEmployee(employee *models.Employee) *Employee {
	return &Employee{
		ID:        employee.ID,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
	}
}

func MultiFromModelToRepoEmployee(employees []models.Employee) []Employee {
	out := make([]Employee, len(employees))

	for i, e := range employees {
		repoEmployee := FromModelToRepoEmployee(&e)
		out[i] = *repoEmployee
	}

	return out
}

func ToModelFromRepoEmployee(employee *Employee) *models.Employee {
	return &models.Employee{
		ID:        employee.ID,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
	}
}

func MultiToModelFromRepoEmployee(employees []Employee) []models.Employee {
	out := make([]models.Employee, len(employees))

	for i, e := range employees {
		modelEmployee := ToModelFromRepoEmployee(&e)
		out[i] = *modelEmployee
	}

	return out
}
