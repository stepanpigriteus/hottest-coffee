package handler

import (
	"encoding/json"
	"fmt"
	"hot/internal/dal"
	"hot/internal/service"
	"hot/models"
	"net/http"
	"strings"
)

type inventoryHandler struct{}

func (i *inventoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	endpoint := strings.Split(r.URL.Path, "/")

	if r.URL.Path == "/inventory/" {
		fmt.Println(r.URL.Path, r.Method)
		switch r.Method {
		case "POST":
			i.postNewItem(w, r)
		case "GET":
			i.getAllItems(w, r)
		default:
			i.undefinedError(w, r)
		}
		return
	}

	if len(endpoint) == 3 && endpoint[0] == "inventory" {
		id := endpoint[2]
		switch r.Method {
		case "GET":
			i.getItemById(w, r, id)
		case "PUT":
			i.putUpdateItemById(w, r, id)
		case "POST":
			i.postNewItem(w, r)
		case "DELETE":

			i.deleteDeleteItemById(w, r, id)
		default:
			i.undefinedError(w, r)
		}
		return
	}

	i.undefinedError(w, r)
}

func (i *inventoryHandler) postNewItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.InventoryItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := models.Error{Message: "Request body does not match json format"}
		json.NewEncoder(w).Encode(response)
		return
	}
	defer r.Body.Close()
	status, msg := service.PostItem(&item)

	w.WriteHeader(status)
	if msg != "" {
		w.WriteHeader(500)

		response := models.Error{Message: msg}
		json.NewEncoder(w).Encode(response)

	}
}

func (i *inventoryHandler) getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	items, status := service.GetItems()
	if status != http.StatusOK {
		w.WriteHeader(status)
		response := models.Error{Message: "Internal server error occurred"}
		json.NewEncoder(w).Encode(response)
		return
	}

	byteValue, err := json.MarshalIndent(items, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		response := models.Error{Message: "Failed to generate json-response"}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(byteValue)

	}
}

func (i *inventoryHandler) getItemById(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")

	inventoryItem, status, msg := service.GetItemById(id)
	if msg != "" {
		w.WriteHeader(status)

		response := models.Error{Message: msg}
		json.NewEncoder(w).Encode(response)
		return
	}
	byteValue, err := json.MarshalIndent(inventoryItem, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		response := models.Error{Message: "Failed to generate json-response"}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(byteValue)

	}
}

func (i *inventoryHandler) putUpdateItemById(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")

	var item models.InventoryItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := models.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(response)

		return
	}
	dall := new(dal.Items)
	err = dall.PutItemById(item, id)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		response := models.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(response)

		return
	}
}

func (i *inventoryHandler) deleteDeleteItemById(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")

	status, msg := service.DeleteItemById(id)
	w.WriteHeader(status)
	if msg != "" {
		w.WriteHeader(500)
		response := models.Error{Message: msg}
		json.NewEncoder(w).Encode(response)

	}
	w.Header().Set("Content-Type", "application/json")
}

func (i *inventoryHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	response := models.Error{Message: "Undefined Error, please check your method or endpoint correctness"}
	json.NewEncoder(w).Encode(response)
}
