package employee

import (
	dto "records/internal/http/DTO"
	"time"
)

type Employee struct {
	id           int       //уникальный ИД
	fullname     string    //фио
	position     string    //должность
	reception_at time.Time //время трудоустройства
}

var id = 0

func NewEmployee(
	fullname string,
	position string,
) *Employee {
	id++
	return &Employee{
		id:           id,
		fullname:     fullname,
		position:     position,
		reception_at: time.Now(),
	}
}

func (e *Employee) ToDTO() *dto.Output_DTO {
	return &dto.Output_DTO{
		ID:           e.id,
		Fullname:     e.fullname,
		Position:     e.position,
		Reception_at: e.reception_at,
	}
}

func NewEmployeeFromDB(id int, fullname, position string, reception_at time.Time) *Employee {
	return &Employee{
		id:           id,
		fullname:     fullname,
		position:     position,
		reception_at: reception_at,
	}
}
