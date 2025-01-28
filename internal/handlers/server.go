package handler

import (
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.Handle("/orders/", &orderHandler{})

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
