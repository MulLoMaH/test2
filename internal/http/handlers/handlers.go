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

func HTTP_write_Error(w http.ResponseWriter, err error, statusCode int) {
	newErr := dto.Error_DTO{}
	newErr = *newErr.NewError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if encodeErr := json.NewEncoder(w).Encode(newErr); encodeErr != nil {
		log.Printf("failed to encode error DTO: %v", encodeErr)
	}
	//http.Error(w, err.Error, http.StatusBadRequest)
}

func (e *Employee_Handlers) Add_employee(w http.ResponseWriter, r *http.Request) {
	new_Employee := dto.Input_DTO{}

	if err := json.NewDecoder(r.Body).Decode(&new_Employee); err != nil {
		HTTP_write_Error(w, err, 400)
		return
	}

	if err := e.repo.Add_employee(new_Employee.Fullname, new_Employee.Position); err != nil {
		HTTP_write_Error(w, err, 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{
		"create": "OK",
	}); err != nil {
		log.Println("error: ", err)
	}

}

func (e *Employee_Handlers) GetAll(w http.ResponseWriter, r *http.Request) {
	allEmployees, err := e.repo.GetAll()
	if err != nil {
		HTTP_write_Error(w, err, 500)
		return
	}

	response := make(map[string]*dto.Output_DTO)
	for fullName, emp := range allEmployees {
		response[fullName] = emp.ToDTO()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("error: ", err)
	}

}

func (e *Employee_Handlers) Delete_Employee(w http.ResponseWriter, r *http.Request) {
	nameStr := r.URL.Query().Get("name")
	if nameStr == "" {
		HTTP_write_Error(w, errors.New("QUERY no delete value"), 400)
		return
	}

	if err := e.repo.Delete_Employee(nameStr); err != nil {
		HTTP_write_Error(w, err, 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{
		"sacced": "success",
	}); err != nil {
		log.Println("error: ", err)
	}

}
