package service

import (
	"fmt"
	"net/http"

	"hot/internal/dal"
	"hot/models"
)

func GetTotalSales() (float64, int, string) {
	aggreg, err := dal.GetAggregations()
	if err != nil {
		return 0, http.StatusInternalServerError, "Error while accessing json files"
	}
	return aggreg.TotalSales, http.StatusOK, ""
}

func GetPopularItems() ([]models.AggregationItem, int, string) {
	aggreg, err := dal.GetItemSales()
	if err != nil {
		return nil, http.StatusInternalServerError, "Error while accessing json files"
	}
	fmt.Println(aggreg)

	return aggreg, http.StatusOK, ""
}
