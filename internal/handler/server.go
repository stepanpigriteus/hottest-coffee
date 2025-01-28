package handler

import (
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.Handle("/orders/", &orderHandler{})
	mux.Handle("/menu/", &menuHandler{})
	mux.Handle("/inventory/", &inventoryHandler{})
	mux.Handle("/reports/", &aggregationHandler{})

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
