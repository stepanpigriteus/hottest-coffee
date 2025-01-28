package handler

import "net/http"

type aggregationHandler struct{}

func (a *aggregationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("asdasdasd"))
}
