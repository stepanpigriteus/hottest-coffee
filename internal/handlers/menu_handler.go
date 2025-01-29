package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"hot/internal/service"
	"hot/models"
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\n\t\"error\": \"Request body does not match json format\"\n}"))
		return
	}
	defer r.Body.Close()

	status, msg := service.PostMenuItem(&menuItem)

	w.WriteHeader(status)
	if msg != "" {
		w.Write([]byte("{\n\t\"error\": \"" + msg + "\"\n}"))
	}
	w.Header().Set("Content-Type", "application/json")
}

func (m *menuHandler) getAllItems(w http.ResponseWriter, r *http.Request) {
	menuItems, status := service.GetMenu()
	if status != http.StatusOK {
		w.WriteHeader(status)
		w.Write([]byte("{\n\t\"error\": \"Internal server occurred\"\n}")) // TODO: Change(?)
		return
	}

	byteValue, err := json.MarshalIndent(menuItems, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\n\t\"error\": \"Failed to generate json-response\"\n}"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(byteValue)
	}
}

func (m *menuHandler) getItemById(w http.ResponseWriter, r *http.Request, id string) {
	menuItem, status, msg := service.GetMenuItemById(id)

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
	status, msg := service.PutUpdateItemBy(&menuItem)
	w.WriteHeader(status)
	if msg != "" {
		w.Write([]byte("{\n\t\"error\": \"" + msg + "\"\n}"))
	}
	w.Header().Set("Content-Type", "application/json")
}

func (m *menuHandler) deleteDeleteItemById(w http.ResponseWriter, r *http.Request, id string) {
	status, msg := service.DeleteMenuItemById(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if msg != "" {
		w.Write([]byte("{\n\t\"error\": \"" + msg + "\"\n}"))
	}
	w.Header().Set("Content-Type", "application/json")
}

func (m *menuHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Undefined Error, please check your method or endpoint correctness"))
}
