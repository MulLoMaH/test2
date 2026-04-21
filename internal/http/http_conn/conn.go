package http_conn

import (
	"records/internal/http/handlers"

	"github.com/gorilla/mux"
)

type Router struct {
	mux *mux.Router
}

func Connect_Server(h *handlers.Employee_Handlers) *Router {
	r := mux.NewRouter()

	r.HandleFunc("/employees", h.Add_employee).Methods("POST")
	r.HandleFunc("/employees/all", h.GetAll).Methods("GET")
	r.HandleFunc("/employees", h.Delete_Employee).Methods("DELETE")

	return &Router{mux: r}
}

func (r *Router) GetMux() *mux.Router {
	return r.mux
}
