package handler

import (
	"encoding/json"
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
	json.NewDecoder(r.Body).Decode(&menuItem)
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (m *menuHandler) getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (m *menuHandler) getItemById(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (m *menuHandler) putUpdateItemById(w http.ResponseWriter, r *http.Request, id string) {
	var menuItem models.MenuItem
	json.NewDecoder(r.Body).Decode(&menuItem)
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (m *menuHandler) deleteDeleteItemById(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	// response:=
	// json.NewEncoder(w).Encode(response)
}

func (m *menuHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Undefined Error, please check your method or endpoint correctness"))
}
