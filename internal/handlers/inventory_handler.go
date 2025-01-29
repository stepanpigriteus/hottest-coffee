package handler

import (
	"encoding/json"
	"fmt"
	"hot/internal/dal"
	"hot/models"
	"net/http"
	"strings"
)

type inventoryHandler struct{}

func (i *inventoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	endpoint := strings.Split(r.URL.Path, "/")
	if r.URL.Path == "/inventory/" {
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
	id := endpoint[2]
	if len(endpoint) == 3 {
		switch r.Method {
		case "GET":
			i.getItemById(w, r, id)
		case "PUT":
			i.putUpdateItemById(w, r, id)
		case "POST":
			i.postNewItem(w, r)
		case "DELETE":
		default:
			i.undefinedError(w, r)
		}
		return
	}
}

func (i *inventoryHandler) postNewItem(w http.ResponseWriter, r *http.Request) {
	var item models.InventoryItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	dall := new(dal.Items)

	err = dall.PostItem(&item)
	if err != nil {
		http.Error(w, "Error while creating inventory item: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (i *inventoryHandler) getAllItems(w http.ResponseWriter, r *http.Request) {
	dall := new(dal.Items)
	item, err := dall.GetItems()
	t, err  := dal.GetInventory()
	fmt.Println(t)
	if err != nil {

		http.Error(w, fmt.Sprint("Items  %s not found"), http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}

func (i *inventoryHandler) getItemById(w http.ResponseWriter, r *http.Request, id string) {
	dall := new(dal.Items)
	item, err := dall.GetItemById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Item with ID %s not found", id), http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
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
	dall := new(dal.Items)
	err := dall.DeleteItemById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("The item %s  was not deleted", id), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (i *inventoryHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Undefined Error, please check your method or endpoint correctness"))
}
