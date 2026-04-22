package interface_employee

import (
	"context"
	"records/internal/employee"
)

type Employee_Repository interface {
	Add_employee(ctx context.Context, fullname string, position string) error
	GetAll(ctx context.Context) (map[string]*employee.Employee, error)
	Delete_Employee(ctx context.Context, fullname string) error
}
