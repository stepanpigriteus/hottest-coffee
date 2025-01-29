package handler

import (
	"net/http"
)

type aggregationHandler struct{}

func (a *aggregationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/reports/total-sales" && r.Method == "GET" {
		a.getTotalSales(w, r)
	} else if r.URL.Path == "/reports/popular-items" && r.Method == "GET" {
		a.getPopularItems(w, r)
	} else {
		a.undefinedError(w, r)
	}
}

func (a *aggregationHandler) getTotalSales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (a *aggregationHandler) getPopularItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (a *aggregationHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Undefined Error, please check your method or endpoint correctness"))
}
