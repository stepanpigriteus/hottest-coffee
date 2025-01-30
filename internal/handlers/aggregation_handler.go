package handler

import (
	"encoding/json"
	"fmt"
	"hot/internal/pkg/config"
	"hot/internal/service"
	"hot/models"
	"net/http"
)

type aggregationHandler struct{}

func (a *aggregationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/reports/total-sales/" && r.Method == "GET" {
		a.getTotalSales(w, r)
	} else if r.URL.Path == "/reports/popular-items" && r.Method == "GET" {
		a.getPopularItems(w, r)
	} else {
		a.undefinedError(w, r)
	}
}

func (a *aggregationHandler) getTotalSales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	config.Logger.Info("request received", "method", r.Method, "url", r.URL)
	totalSales, status, errStr := service.GetTotalSales()
	w.WriteHeader(status)
	if status != http.StatusOK {
		w.WriteHeader(500)

		response := models.Error{Message: errStr}
		json.NewEncoder(w).Encode(response)

	} else {
		response := fmt.Sprintf("{\n\t\"total_sales\": %v\n}", totalSales)
		w.Write([]byte(response))

	}
}

func (a *aggregationHandler) getPopularItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	config.Logger.Info("request received", "method", r.Method, "url", r.URL)
	popularItems, status, errStr := service.GetPopularItems()
	w.WriteHeader(status)
	if status != http.StatusOK {
		w.WriteHeader(500)

		response := models.Error{Message: errStr}
		json.NewEncoder(w).Encode(response)
	} else {
		byteValue, err := json.MarshalIndent(popularItems, "", "\t")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := models.Error{Message: "Error while writing response"}
			json.NewEncoder(w).Encode(response)
		} else {
			json.NewEncoder(w).Encode(byteValue)
		}
	}
}

func (a *aggregationHandler) undefinedError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := models.Error{Message: "Undefined Error, please check your method or endpoint correctness"}
	json.NewEncoder(w).Encode(response)
}
