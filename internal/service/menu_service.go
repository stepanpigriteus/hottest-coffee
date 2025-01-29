package service

import (
	"hot/internal/dal"
	"hot/models"
	"net/http"
)

func GetMenuItemById(id string) (models.MenuItem, int, string) {
	var menuRepo dal.MenuInterface = new(dal.MenuItems)
	menuItem, err := menuRepo.GetMenuItemById(id)
	if err != nil {
		if err == models.ErrMenuItemNotFound {
			return models.MenuItem{}, http.StatusNotFound, err.Error()
		}
		return models.MenuItem{}, http.StatusInternalServerError, ""
	}

	return menuItem, http.StatusOK, ""
}
