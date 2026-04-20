package main

import (
	"errors"
	"log"
	"net/http"
	"records/internal/http/http_conn"
)

func main() {
	if err := http.ListenAndServe(":8080", http_conn.Connect_Server().GetMux()); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("normal started err")
		} else {
			log.Fatal(err)
			return
		}
	}

}
