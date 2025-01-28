package handler

import (
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
	if len(endpoint) == 3 {
		switch method {
		case "GET":
			o.getOrderById(w, r)
		case "PUT":
			o.putUpdateOrder(w, r)
		case "DELETE":
			o.deleteDeleteOrder(w, r)
		default:
			o.undefinedError(w, r)
		}
		return
	}
	if len(endpoint) == 4 && r.Method == "POST" && endpoint[3] == "close" {
		o.postCloseOrder(w, r)
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

func (o *orderHandler) getOrderById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getOrderById"))
}

func (o *orderHandler) putUpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("putUpdateOrder"))
}

func (o *orderHandler) deleteDeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleteDeleteOrder"))
}

func (o *orderHandler) postCloseOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("postCloseOrder"))
}

func (o *orderHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Undefined Error, please check your method or endpoint correctness"))
}
