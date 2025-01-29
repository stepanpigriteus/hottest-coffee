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

type menuHandler struct{}

func (m *menuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	endpoint := strings.Split(r.URL.Path, "/")

	if r.URL.Path == "/menu/" {
		switch r.Method {
		case "POST":
			m.postNewItem(w, r)
		case "GET":
			m.getAllItems(w, r)
		default:
			m.undefinedError(w, r)
		}
		return
	}
	id := endpoint[2]
	if len(endpoint) == 3 {
		switch r.Method {
		case "GET":
			m.getItemById(w, r, id)
		case "PUT":
			m.putUpdateItemById(w, r, id)
		case "DELETE":
			m.deleteDeleteItemById(w, r, id)
		default:
			m.undefinedError(w, r)
		}
		return
	}

	m.undefinedError(w, r)
}

func (m *menuHandler) postNewItem(w http.ResponseWriter, r *http.Request) {
	var menuItem models.MenuItem
	err := json.NewDecoder(r.Body).Decode(&menuItem)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	dall := new(dal.MenuItems)
	err = dall.PostMenuItem(&menuItem)
	if err != nil {
		http.Error(w, "Error while creating order: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (m *menuHandler) getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dall := new(dal.MenuItems)

	i, err := dall.GetMenuItems()
	if err != nil {
		http.Error(w, fmt.Sprint("Orders  %s not found"), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(i); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (m *menuHandler) getItemById(w http.ResponseWriter, r *http.Request, id string) {
	menuItem, status, msg := service.GetMenuItemById(id)
	fmt.Println(menuItem, status, msg)

	if msg != "" {
		w.WriteHeader(status)
		w.Write([]byte("{\n\t\"error\": \"" + msg + "\"\n}"))
		return
	}

	byteValue, err := json.MarshalIndent(menuItem, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\n\t\"error\": \"Failed to generate json-response\"\n}"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(byteValue)
	}

	w.Header().Set("Content-Type", "application/json")
}

func (m *menuHandler) putUpdateItemById(w http.ResponseWriter, r *http.Request, id string) {
	var menuItem models.MenuItem
	json.NewDecoder(r.Body).Decode(&menuItem)
	dall := new(dal.MenuItems)
	err := dall.PutMenuItemById(&menuItem, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("MenuItem with ID %s not found", id), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "Error while creating/update menuItem: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (m *menuHandler) deleteDeleteItemById(w http.ResponseWriter, r *http.Request, id string) {
	dall := new(dal.MenuItems)
	err := dall.DeleteMenuItemById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("The menuItem %s  was not deleted", id), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (m *menuHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Undefined Error, please check your method or endpoint correctness"))
}
