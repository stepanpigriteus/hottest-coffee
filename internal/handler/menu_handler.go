package server

import "net/http"

type menuHandler struct{}

func (m *menuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
