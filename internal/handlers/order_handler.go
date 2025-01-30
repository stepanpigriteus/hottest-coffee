package handler

import (
	"encoding/json"
	"hot/internal/service"
	"hot/models"
	"net/http"
	"strings"
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
	w.Header().Set("Content-Type", "application/json")
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		response := models.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	status, msg := service.PostOrder(&order)

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
	w.Header().Set("Content-Type", "application/json")
	orders, status := service.GetOrders()

	w.WriteHeader(status)
	if status != http.StatusOK {
		w.WriteHeader(500)

		response := models.Error{Message: "Internal server error occurred"}
		json.NewEncoder(w).Encode(response)
		return
	}

	byteValue, err := json.MarshalIndent(orders, "", "\t")
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		response := models.Error{Message: "Failed to generate json-response"}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(byteValue)

	}
}

func (o *orderHandler) getOrderById(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	order, status, msg := service.GetOrderById(id)

	if msg != "" {
		w.WriteHeader(status)

		response := models.Error{Message: msg}
		json.NewEncoder(w).Encode(response)
		return
	}

	byteValue, err := json.MarshalIndent(order, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		response := models.Error{Message: "Failed to generate json-response"}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(byteValue)
	}
}

func (o *orderHandler) putUpdateOrder(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)
	status, msg := service.PutOrderById(&order, id)
	w.WriteHeader(status)
	if msg != "" {
		w.WriteHeader(500)

		response := models.Error{Message: msg}
		json.NewEncoder(w).Encode(response)
	}
}

func (o *orderHandler) deleteDeleteOrder(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	status, msg := service.DeleteOrderById(id)
	w.WriteHeader(status)
	if msg != "" {

		w.WriteHeader(500)

		response := models.Error{Message: msg}
		json.NewEncoder(w).Encode(response)
	}
}

func (o *orderHandler) postCloseOrder(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	status, msg := service.CloseOrder(id)
	response := Response{Message: msg}

	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(response); err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		response2 := models.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(response2)
	}
}

func (o *orderHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	response := models.Error{Message: "Undefined Error, please check your method or endpoint correctness"}
	json.NewEncoder(w).Encode(response)
}
