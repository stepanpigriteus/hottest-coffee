package handler

import (
	"encoding/json"
	"hot/models"
	"net/http"
	"strings"
)

type orderHandler struct{}

func (o *orderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	endpoint := strings.Split(r.URL.Path, "/")
	if r.URL.Path == "/orders/" {
		switch r.Method {
		case "POST":
			o.postCreateOrder(w, r)
		case "GET":
			o.getAllOrders(w, r)
		default:
			o.undefinedError(w, r)
		}
		return
	}
	id := endpoint[2]
	if len(endpoint) == 3 {
		switch r.Method {
		case "GET":
			o.getOrderById(w, r, id)
		case "PUT":
			o.putUpdateOrder(w, r, id)
		case "DELETE":
			o.deleteDeleteOrder(w, r, id)
		default:
			o.undefinedError(w, r)
		}
		return
	}
	if len(endpoint) == 4 && r.Method == "POST" && endpoint[3] == "close" {
		o.postCloseOrder(w, r, id)
		return
	}
	o.undefinedError(w, r)
}

func (o *orderHandler) postCreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (o *orderHandler) getAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (o *orderHandler) getOrderById(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (o *orderHandler) putUpdateOrder(w http.ResponseWriter, r *http.Request, id string) {
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (o *orderHandler) deleteDeleteOrder(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (o *orderHandler) postCloseOrder(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (o *orderHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Undefined Error, please check your method or endpoint correctness"))
}
