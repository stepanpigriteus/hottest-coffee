package handler

import "net/http"

type AgregationHandler struct{}

func (h *AgregationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my agregß page"))
}
