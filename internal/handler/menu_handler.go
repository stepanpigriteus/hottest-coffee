package handler

import (
	"encoding/json"
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
	w.Header().Set("Content-Type", "application/json")

	var menuItem models.MenuItem
	err := json.NewDecoder(r.Body).Decode(&menuItem)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		response := models.Error{Message: "Request body does not match json format"}

		json.NewEncoder(w).Encode(response)

		return
	}
	defer r.Body.Close()

	status, msg := service.PostMenuItem(&menuItem)

	w.WriteHeader(status)
	if msg != "" {
		w.WriteHeader(http.StatusBadRequest)
		response := models.Error{Message: msg}
		json.NewEncoder(w).Encode(response)
	}
}

func (m *menuHandler) getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	menuItems, status := service.GetMenu()
	if status != http.StatusOK {
		w.WriteHeader(status)
		response := models.Error{Message: "Internal Server Error Occured"}

		json.NewEncoder(w).Encode(response)
		return
	}

	byteValue, err := json.MarshalIndent(menuItems, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := models.Error{Message: "Failed to generate json-response"}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(byteValue)

	}
}

func (m *menuHandler) getItemById(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	menuItem, status, msg := service.GetMenuItemById(id)

	if msg != "" {
		w.WriteHeader(status)

		response := models.Error{Message: msg}
		json.NewEncoder(w).Encode(response)
		return
	}

	byteValue, err := json.MarshalIndent(menuItem, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		response := models.Error{Message: "Failed to generate json-response"}
		json.NewEncoder(w).Encode(response)

	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(byteValue)
	}
}

func (m *menuHandler) putUpdateItemById(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	var menuItem models.MenuItem
	json.NewDecoder(r.Body).Decode(&menuItem)
	status, msg := service.PutUpdateItemBy(&menuItem)
	w.WriteHeader(status)
	if msg != "" {
		w.WriteHeader(500)

		response := models.Error{Message: msg}
		json.NewEncoder(w).Encode(response)
	}
}

func (m *menuHandler) deleteDeleteItemById(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	status, msg := service.DeleteMenuItemById(id)
	w.WriteHeader(status)
	if msg != "" {
		w.WriteHeader(500)

		response := models.Error{Message: msg}
		json.NewEncoder(w).Encode(response)
	}
}

func (m *menuHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	response := models.Error{Message: "Undefined Error, please check your method or endpoint correctness"}
	json.NewEncoder(w).Encode(response)
}
