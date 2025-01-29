package handler

import (
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.Handle("/orders/", loggingMiddleware(&orderHandler{}))
	mux.Handle("/menu/", loggingMiddleware(&menuHandler{}))
	mux.Handle("/inventory/", loggingMiddleware(&inventoryHandler{}))
	mux.Handle("/reports/", loggingMiddleware(&aggregationHandler{}))
	err := http.ListenAndServe(":8200", mux)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
