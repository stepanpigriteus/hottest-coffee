package handler

import "net/http"

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("{\n\t\"error\": \"Wrong method and/or URL\"\n}"))
}
