package server

import "net/http"

type orderHandler struct{}

func (o *orderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("asdasdasd"))
}
