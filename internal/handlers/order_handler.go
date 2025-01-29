package handler

import (
	"encoding/json"
	"fmt"
	"hot/internal/dal"
	"hot/models"
	"net/http"
	"strings"
)

type orderHandler struct{}

func (o *orderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	endpoint := strings.Split(r.URL.Path, "/")
	method := r.Method
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
		switch method {
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
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {

		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	dall := new(dal.Orders)
	err = dall.PostOrder(&order)
	if err != nil {
		http.Error(w, "Error while creating order: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (o *orderHandler) getAllOrders(w http.ResponseWriter, r *http.Request) {
	dall := new(dal.Orders)
	i, err := dall.GetOrders()
	if err != nil {

		http.Error(w, fmt.Sprint("Orders  %s not found"), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(i); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (o *orderHandler) getOrderById(w http.ResponseWriter, r *http.Request, id string) {
	dall := new(dal.Orders)

	i, err := dall.GetOrderById(id)
	if err != nil {

		http.Error(w, fmt.Sprintf("Order with ID %s not found", id), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(i); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (o *orderHandler) putUpdateOrder(w http.ResponseWriter, r *http.Request, id string) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {

		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	dall := new(dal.Orders)
	err = dall.PutUpdate(&order, id)
	if err != nil {
		http.Error(w, "Error while creating order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (o *orderHandler) deleteDeleteOrder(w http.ResponseWriter, r *http.Request, id string) {
	dall := new(dal.Orders)
	err := dall.DeleteOrderById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("The order %s  was not deleted", id), http.StatusBadRequest)
		return
	}
}

func (o *orderHandler) postCloseOrder(w http.ResponseWriter, r *http.Request, id string) {
	dall := new(dal.Orders)
	i, err := dall.CloseOrder(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("The order %s  was not deleted", id), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(i); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (o *orderHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Undefined Error, please check your method or endpoint correctness"))
}
