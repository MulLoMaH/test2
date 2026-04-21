package dto

import "time"

//data transfer objects

// входящий запрос
type Input_DTO struct {
	Fullname string `json:"Full_Name"`
	Position string `json:"Position"`
}

// исходящее тело ответа
type Output_DTO struct {
	ID           int       `json:"id"`
	Fullname     string    `json:"full_Name"`
	Position     string    `json:"position"`
	Reception_at time.Time `json:"reception_at"`
}

// структура для передачи ошибок
type Error_DTO struct {
	Err      error     `json:"error"`
	Time_Err time.Time `json:"time_Err"`
}

func (e Error_DTO) NewError(err error) *Error_DTO {
	return &Error_DTO{
		Err:      err,
		Time_Err: time.Now(),
	}
}
