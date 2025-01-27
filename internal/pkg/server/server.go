package server

import (
	"net/http"

	"hot/internal/handler"
)

func Start(port int, dir string) {
	mux := http.NewServeMux()

	mux.Handle("/orders", &handler.OrdersHandler{})
	mux.Handle("/menu", &handler.MenuHandler{})
	mux.Handle("/inventory", &handler.InventoryHandler{})
	mux.Handle("/reports/total-sales", &handler.AgregationHandler{})
	mux.Handle("/reports/popular-items", &handler.AgregationHandler{})
	mux.HandleFunc("/", handler.BadRequest)
	http.ListenAndServe(":8080", mux)
}
