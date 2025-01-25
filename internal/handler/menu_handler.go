package handler

import "net/http"

type MenuHandler struct{}

func (h *MenuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my men page"))
}
