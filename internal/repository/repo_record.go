package repository

import (
	"errors"
	"records/internal/employee"
	"sync"
)

type Employees struct {
	employees map[string]*employee.Employee
	mtx       sync.RWMutex
}

func NewEmployees() *Employees {
	return &Employees{
		employees: make(map[string]*employee.Employee),
	}
}

func (e *Employees) Add_employee(fullname string, position string) error {
	if err := validate(fullname, position); err != nil {
		return err
	}

	newEmployees := employee.NewEmployee(fullname, position)

	e.mtx.Lock()
	defer e.mtx.Unlock()
	e.employees[fullname] = newEmployees

	return nil
}

func (e *Employees) GetAll() (map[string]*employee.Employee, error) {
	all_employee := make(map[string]*employee.Employee)

	e.mtx.RLock()
	defer e.mtx.RUnlock()

	if len(e.employees) == 0 {
		return all_employee, errors.New("No Content")
	}
	for k, v := range e.employees {
		all_employee[k] = v
	}

	return all_employee, nil
}

func (e *Employees) Delete_Employee(fullname string) error {
	e.mtx.Lock()
	defer e.mtx.Unlock()
	_, ok := e.employees[fullname]
	if !ok {
		return errors.New("Employee not found")
	}

	delete(e.employees, fullname)

	return nil
}

func validate(fullname string, position string) error {
	if fullname == "" {
		return errors.New("Fullname in Empty")
	}

	if position == "" {
		return errors.New("Position in Empty")
	}

	return nil
}
