package handler

import (
	"encoding/json"
	"fmt"
	"hot/models"
	"log"
	"net/http"
)

func StartServer(addr int) {
	mux := http.NewServeMux()
	mux.Handle("/", loggingMiddleware(&handleDef{}))
	mux.Handle("/orders/", loggingMiddleware(&orderHandler{}))
	mux.Handle("/menu/", loggingMiddleware(&menuHandler{}))
	mux.Handle("/inventory/", loggingMiddleware(&inventoryHandler{}))
	mux.Handle("/reports/", loggingMiddleware(&aggregationHandler{}))
	port := fmt.Sprintf(":%d", addr)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

type handleDef struct{}

func (h *handleDef) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	response := models.Error{Message: "Undefined Error, please check your method or endpoint correctness"}
	json.NewEncoder(w).Encode(response)
}
