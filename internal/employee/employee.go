package employee

import "time"

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
