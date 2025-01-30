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
	var item models.InventoryItem

	status, msg := service.PostItem(&item)

	w.WriteHeader(status)
	if msg != "" {
		w.Write([]byte("{\n\t\"error\": \"" + msg + "\"\n}"))
	}
	w.Header().Set("Content-Type", "application/json")
}

func (i *inventoryHandler) getAllItems(w http.ResponseWriter, r *http.Request) {
	items, status := service.GetItems()
	if status != http.StatusOK {
		w.WriteHeader(status)
		w.Write([]byte("{\n\t\"error\": \"Internal server occurred\"\n}"))
		return
	}

	byteValue, err := json.MarshalIndent(items, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\n\t\"error\": \"Failed to generate json-response\"\n}"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(byteValue)
	}
}

func (i *inventoryHandler) getItemById(w http.ResponseWriter, r *http.Request, id string) {
	inventoryItem, status, msg := service.GetItemById(id)
	if msg != "" {
		w.WriteHeader(status)
		w.Write([]byte("{\n\t\"error\": \"" + msg + "\"\n}"))
		return
	}
	byteValue, err := json.MarshalIndent(inventoryItem, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\n\t\"error\": \"Failed to generate json-response\"\n}"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(byteValue)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (i *inventoryHandler) putUpdateItemById(w http.ResponseWriter, r *http.Request, id string) {
	var item models.InventoryItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	dall := new(dal.Items)
	err = dall.PutItemById(item, id)
	if err != nil {
		http.Error(w, "Error while creating item: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (i *inventoryHandler) deleteDeleteItemById(w http.ResponseWriter, r *http.Request, id string) {
	status, msg := service.DeleteItemById(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if msg != "" {
		w.Write([]byte("{\n\t\"error\": \"" + msg + "\"\n}"))
	}
	w.Header().Set("Content-Type", "application/json")
}

func (i *inventoryHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Undefined Error, please check your method or endpoint correctness"))
}
