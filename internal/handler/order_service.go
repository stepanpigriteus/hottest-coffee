package handler

import (
	"encoding/json"
	"net/http"
)

type OrdersHandler struct{}

func (h *OrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "",
	}

	switch r.Method {
	case "GET":
		response["message"] = "get"
	case "POST":
		response["message"] = "post"
	case "PUT":
		response["message"] = "put"
	case "DELETE":
		response["message"] = "delete"
	default:
		response["message"] = "method isn't allowed"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
}
