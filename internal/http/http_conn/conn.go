package http_conn

import (
	"github.com/gorilla/mux"
)

type Router struct {
	mux *mux.Router
}

func Connect_Server() *Router {
	r := mux.NewRouter()

	r.HandleFunc("employees").Methods("POST")
	r.HandleFunc("employees").Methods("GET")
	r.HandleFunc("employees").Methods("DELETE")

	return &Router{mux: r}
}

func (r *Router) GetMux() *mux.Router {
	return r.mux
}
