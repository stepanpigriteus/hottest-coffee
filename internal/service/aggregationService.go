package service

import (
	"net/http"
	"sort"

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

func GetPopularItems() ([]models.MenuItem, int, string) {
	aggreg, err := dal.GetAggregations()
	if err != nil {
		return nil, http.StatusInternalServerError, "Error while accessing json files"
	}
	sort.Slice(aggreg.ItemSales, func(i, j int) bool {
		return aggreg.ItemSales[i].QuantSold > aggreg.ItemSales[j].QuantSold
	})
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
