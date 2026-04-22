package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"records/internal/http/handlers"
	"records/internal/http/http_conn"
	"records/internal/repository/database/db_conn"
	dbinteraction "records/internal/repository/database/db_interaction"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ctx := context.Background()

	pool, err := db_conn.ConnDB(ctx)
	if err != nil {
		log.Fatal("error: ", err)
	}

	repo := dbinteraction.NewPostgresRepo(pool)
	handler := handlers.New_Employee_Handlers(repo)

	r := http_conn.Connect_Server(handler)

	fmt.Println("start HTTP server")
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
			log.Println("good start server")
		}
	}()

	wg.Wait()

}
