package server

import (
	"hot/internal/handler"
	"log"
	"net/http"
)

func Start(port int, dir string) {
	mux := http.NewServeMux()

	mux.Handle("/orders", &handler.OrdersHandler{})
	mux.Handle("/menu", &handler.MenuHandler{})
	mux.Handle("/inventory", &handler.InventoryHandler{})
	mux.Handle("/reports/total-sales", &handler.AgregationHandler{})
	mux.Handle("/reports/popular-items", &handler.AgregationHandler{})
	mux.HandleFunc("/", handler.BadRequest)
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
