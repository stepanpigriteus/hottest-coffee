package service

import (
	"hot/internal/dal"
	"hot/models"
	"net/http"
	"sort"
)

func GetTotalSales() (float64, int, string) {
	aggreg, err := dal.GetAggregations()
	if err != nil {
		return 0, http.StatusInternalServerError, "Error while accessing json files"
	}
	return aggreg.TotalSales, http.StatusOK, ""
}

func GetPopularItems() ([]models.MenuItem, int, string) {
	aggreg, err := dal.GetAggregations()
	if err != nil {
		return nil, http.StatusInternalServerError, "Error while accessing json files"
	}
	if len(aggreg.ItemSales) > 0 {
		sort.Slice(aggreg.ItemSales, func(i, j int) bool {
			return aggreg.ItemSales[i].QuantSold > aggreg.ItemSales[j].QuantSold
		})
	}
	var populars []models.MenuItem
	var menuRepo dal.MenuInterface = new(dal.MenuItems)
	for _, v := range aggreg.ItemSales[:min(3, len(aggreg.ItemSales))] {
		menuItem, err := menuRepo.GetMenuItemById(v.ID)
		if err != nil {
			return nil, http.StatusInternalServerError, err.Error()
		}
		populars = append(populars, menuItem)
	}
	return populars, http.StatusOK, ""
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
