package server

import "net/http"

func StartServer() {
	mux := http.NewServeMux()
	mux.Handle("/orders/", &orderHandler{})
	mux.Handle("/menu/", &menuHandler{})
	mux.Handle("/inventory/", &inventoryHandler{})
	mux.Handle("/reports/", &aggregationHandler{})

	http.ListenAndServe(":8000", mux)
}
