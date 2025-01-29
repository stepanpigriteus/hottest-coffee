package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"hot/internal/service"
	"hot/models"
)

type orderHandler struct{}

type Response struct {
	Message string `json:"message"`
}

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

	status, msg := service.PostOrder(&order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := make(map[string]string)
	if status == http.StatusCreated {
		response["message"] = msg
	} else {
		response["error"] = msg
	}

	json.NewEncoder(w).Encode(response)
}

func (o *orderHandler) getAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, status := service.GetOrders()

	w.WriteHeader(status)
	if status != http.StatusOK {

		w.Write([]byte("{\n\t\"error\": \"Internal server occurred\"\n}")) // TODO: Change(?)
		return
	}

	byteValue, err := json.MarshalIndent(orders, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\n\t\"error\": \"Failed to generate json-response\"\n}"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(byteValue)
	}
}

func (o *orderHandler) getOrderById(w http.ResponseWriter, r *http.Request, id string) {
	order, status, msg := service.GetOrderById(id)

	if msg != "" {
		w.WriteHeader(status)
		w.Write([]byte("{\n\t\"error\": \"" + msg + "\"\n}"))
		return
	}

	byteValue, err := json.MarshalIndent(order, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\n\t\"error\": \"Failed to generate json-response\"\n}"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(byteValue)
	}

	w.Header().Set("Content-Type", "application/json")
}

func (o *orderHandler) putUpdateOrder(w http.ResponseWriter, r *http.Request, id string) {
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)
	status, msg := service.PutOrderById(&order, id)
	w.WriteHeader(status)
	if msg != "" {
		w.Write([]byte("{\n\t\"" + msg + "\"\n}"))
	}
	w.Header().Set("Content-Type", "application/json")
}

func (o *orderHandler) deleteDeleteOrder(w http.ResponseWriter, r *http.Request, id string) {
	status, msg := service.DeleteOrderById(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if msg != "" {
		w.Write([]byte("{\n\t\"error\": \"" + msg + "\"\n}"))
	}
	w.Header().Set("Content-Type", "application/json")
}

func (o *orderHandler) postCloseOrder(w http.ResponseWriter, r *http.Request, id string) {
	status, msg := service.CloseOrder(id)
	response := Response{Message: msg}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (o *orderHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Undefined Error, please check your method or endpoint correctness"))
}
