package main

import (
	"errors"
	"log"
	"net/http"
	"records/internal/http/handlers"
	"records/internal/http/http_conn"
	"records/internal/repository"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	repo := repository.NewEmployees()
	handler := handlers.New_Employee_Handlers(repo)

	r := http_conn.Connect_Server(handler)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(":8080", r.GetMux()); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Println("normal started err")
			} else {
				log.Fatal(err)
				return
			}
		}
	}()

	wg.Wait()

}
