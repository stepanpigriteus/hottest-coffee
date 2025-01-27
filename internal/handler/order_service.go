package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type OrdersHandler struct{}

func (h *OrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
		return
	}
	log.Printf("Received %s request\n", r.Method)

	var response interface{}

	switch r.Method {
	case "GET":
		response = []map[string]interface{}{
			{
				"order_id":      "order123",
				"customer_name": "Alice Smith",
				"items": []map[string]interface{}{
					{"product_id": "latte", "quantity": 2},
					{"product_id": "muffin", "quantity": 1},
				},
				"status":     "open",
				"created_at": "2023-10-01T09:00:00Z",
			},
			{
				"order_id":      "order124",
				"customer_name": "Bob Johnson",
				"items": []map[string]interface{}{
					{"product_id": "espresso", "quantity": 1},
				},
				"status":     "closed",
				"created_at": "2023-10-01T09:30:00Z",
			},
		}
	case "POST":
		parts := strings.Split(r.URL.Path, "/")
		orderId := parts[2]
		log.Printf("Received POST request for orderId: %s", orderId)
		if len(parts) < 4 || parts[1] != "orders" {
			http.Error(w, "Invalid URL path", http.StatusBadRequest)
			return
		}

		if len(parts) >= 4 && parts[3] == "close" {
			// Закрываем заказ
			fmt.Printf("Closing order: %s\n", orderId)
			response = map[string]interface{}{
				"message":  "order closed",
				"order_id": orderId,
			}
		} else {
			http.Error(w, "Invalid URL path for POST request", http.StatusBadRequest)
			return
		}

	case "PUT":

	case "DELETE":
		response = map[string]interface{}{"message": "delete"}
	default:
		response = map[string]interface{}{"message": "method isn't allowed"}
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	// Заголовок Content-Type устанавливаем до отправки данных
	w.Header().Set("Content-Type", "application/json")

	// Кодируем в JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetOrders(w http.ResponseWriter, r *http.Request) {}
