package handler

import (
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.Handle("/orders/", &orderHandler{})
	mux.Handle("/menu/", &menuHandler{})
	err := http.ListenAndServe(":8200", mux)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
