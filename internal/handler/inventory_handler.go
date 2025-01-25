package handler

import "net/http"

type InventoryHandler struct{}

func (h *InventoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my innv page"))
}
