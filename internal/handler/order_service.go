package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type OrdersHandler struct{}

func (h *OrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Разрешить доступ с любого домена
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		log.Println("Handling CORS preflight request")
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
		response = map[string]interface{}{"message": "post"}
	case "PUT":
		response = map[string]interface{}{"message": "put"}
	case "DELETE":
		response = map[string]interface{}{"message": "delete"}
	default:
		response = map[string]interface{}{"message": "method isn't allowed"}
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetOrders(w http.ResponseWriter, r *http.Request) {}
