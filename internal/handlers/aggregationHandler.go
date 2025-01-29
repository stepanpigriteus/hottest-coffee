package handler

import (
	"encoding/json"
	"fmt"
	"hot/internal/pkg/config"
	"hot/internal/service"
	"net/http"
)

type aggregationHandler struct{}

func (o *aggregationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func GetTotalSales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	config.Logger.Info("request received", "method", r.Method, "url", r.URL)
	totalSales, status, errStr := service.GetTotalSales()
	w.WriteHeader(status)
	if status != http.StatusOK {
		w.Write([]byte("{\n\t\"error\": \"" + errStr + "\"\n}"))
	} else {
		w.Write([]byte(fmt.Sprintf("{\n\t\"total_sales\": %v\n}", totalSales)))
	}
}

func GetPopularItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	config.Logger.Info("request received", "method", r.Method, "url", r.URL)
	popularItems, status, errStr := service.GetPopularItems()
	w.WriteHeader(status)
	if status != http.StatusOK {
		w.Write([]byte("{\n\t\"error\": \"" + errStr + "\"\n}"))
	} else {
		byteValue, err := json.MarshalIndent(popularItems, "", "\t")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\n\t\"error\": \"Error while writing response\"\n}"))
		} else {
			w.Write(byteValue)
		}
	}
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("{\n\t\"error\": \"Wrong method and/or URL\"\n}"))
}
