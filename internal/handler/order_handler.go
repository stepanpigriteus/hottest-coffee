package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

type orderHandler struct{}

func (o *orderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	endpoint := strings.Split(r.URL.Path, "/")
	method := r.Method
	if r.URL.Path == "/orders/" {
		switch r.Method {
		case "POST":
			o.postCreateOrder(w, r)
		case "GET":
			o.getAllOrders(w, r)
		default:
			o.undefinedError(w, r)
		}
		return
	}
	id := endpoint[2]
	if len(endpoint) == 3 {
		switch method {
		case "GET":
			o.getOrderById(w, r, id)
		case "PUT":
			o.putUpdateOrder(w, r, id)
		case "DELETE":
			o.deleteDeleteOrder(w, r, id)
		default:
			o.undefinedError(w, r)
		}
		return
	}
	if len(endpoint) == 4 && r.Method == "POST" && endpoint[3] == "close" {
		o.postCloseOrder(w, r, id)
		return
	}
	o.undefinedError(w, r)
}

func (o *orderHandler) postCreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("postCreateOrder"))
}

func (o *orderHandler) getAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getAllOrders"))
}

func (o *orderHandler) getOrderById(w http.ResponseWriter, r *http.Request, id string) {
	response, err := myGet()
	if err != nil {
		w.Write([]byte("internal server error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.Write([]byte("error encoding json"))
		return
	}
}

func (o *orderHandler) putUpdateOrder(w http.ResponseWriter, r *http.Request, id string) {
	w.Write([]byte("putUpdateOrder"))
}

func (o *orderHandler) deleteDeleteOrder(w http.ResponseWriter, r *http.Request, id string) {
	w.Write([]byte("deleteDeleteOrder"))
}

func (o *orderHandler) postCloseOrder(w http.ResponseWriter, r *http.Request, id string) {
	w.Write([]byte("postCloseOrder"))
}

func (o *orderHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Undefined Error, please check your method or endpoint correctness"))
}

func myGet() (string, error) {
	return "response", nil
}
