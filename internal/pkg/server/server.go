package server

import (
	"hot/internal/handler"
	"log"
	"net/http"
)

func Start(port int, dir string) {
	mux := http.NewServeMux()
	mux.Handle("/", &homeHandler{})
	mux.Handle("/orders", &handler.OrdersHandler{})
	mux.Handle("/menu", &handler.MenuHandler{})
	mux.Handle("/inventory", &handler.InventoryHandler{})
	mux.Handle("/reports/total-sales", &handler.AgregationHandler{})
	mux.Handle("/reports/popular-items", &handler.AgregationHandler{})
	err := http.ListenAndServe(":8083", mux)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
