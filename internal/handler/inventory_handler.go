package handler

import (
	"encoding/json"
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
		case "DELETE":
		default:
			i.undefinedError(w, r)
		}
		return
	}
}

func (i *inventoryHandler) postNewItem(w http.ResponseWriter, r *http.Request) {
	var item models.InventoryItem
	json.NewDecoder(r.Body).Decode(&item)
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (i *inventoryHandler) getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (i *inventoryHandler) getItemById(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (i *inventoryHandler) putUpdateItemById(w http.ResponseWriter, r *http.Request, id string) {
	var item models.InventoryItem
	json.NewDecoder(r.Body).Decode(&item)
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (i *inventoryHandler) deleteDeleteItemById(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (i *inventoryHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Undefined Error, please check your method or endpoint correctness"))
}
