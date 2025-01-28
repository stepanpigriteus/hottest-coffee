package handler

import "net/http"

type inventoryHandler struct{}

func (i *inventoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("asdasdasd"))
}
