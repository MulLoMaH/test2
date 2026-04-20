package interface_employee

import "records/internal/employee"

type Employee_Repository interface {
	Add_employee(fullname string, position string) error
	GetAll() (map[string]*employee.Employee, error)
	Delete_Employee(fullname string) error
}
