package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	dto "records/internal/http/DTO"
	"records/internal/repository"
)

type Employee_Handlers struct {
	repo *repository.Employees
}

func New_Employee_Handlers(repo *repository.Employees) *Employee_Handlers {
	return &Employee_Handlers{
		repo: repo,
	}
}

func (e *Employee_Handlers) Add_employee(w http.ResponseWriter, r *http.Request) {

}

func (e *Employee_Handlers) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (e *Employee_Handlers) Delete_Employee(w http.ResponseWriter, r *http.Request) {
	nameStr := r.URL.Query().Get("name")
	if nameStr == "" {
		errDto := dto.Error_DTO{}
		errDto = *errDto.NewError(errors.New("QUERY no delete value"))
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errDto); err != nil {
			errDto = *errDto.NewError(err)
			log.Fatal(errDto)
		}
	}

	if err := e.repo.Delete_Employee(nameStr); err != nil {
		errDto := dto.Error_DTO{}
		w.WriteHeader(http.StatusBadRequest)
		errDto = *errDto.NewError(err)
		if err := json.NewEncoder(w).Encode(errDto); err != nil {
			errDto = *errDto.NewError(err)
			log.Fatal(errDto)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	if err := json.NewEncoder(w).Encode(map[string]string{
		"sacced": "successful",
	}); err != nil {
		errDto := dto.Error_DTO{}
		errDto = *errDto.NewError(err)
		log.Fatal(errDto)
	}

}
