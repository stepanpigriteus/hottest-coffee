package server

import "net/http"

func Start(port int, dir string) {
	mux := http.NewServeMux()
	mux.Handle("/", &homeHandler{})
	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
