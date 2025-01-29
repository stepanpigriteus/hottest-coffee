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

func GetMenu() ([]models.MenuItem, int) {
	var menuRepo dal.MenuInterface = new(dal.MenuItems)
	menuItems, err := menuRepo.GetMenuItems()
	if err != nil {
		return make([]models.MenuItem, 0), http.StatusInternalServerError
	}

	return menuItems, http.StatusOK
}

func PutUpdateItemBy(menuItem *models.MenuItem) (int, string) {
	if !ValidItem(menuItem) {
		return http.StatusBadRequest, models.ErrInvalidMenuItem.Error()
	}
}

func ValidItem(menuItem *models.MenuItem) bool {
	if menuItem == nil {
		return false
	}

	if menuItem.Description == "" || menuItem.ID == "" || menuItem.Name == "" || len(menuItem.Ingredients) <= 0 || menuItem.Price <= 0 {
		return false
	}

	for _, ingr := range menuItem.Ingredients {
		if ingr.IngredientID == "" || ingr.Quantity <= 0 {
			return false
		}
	}

	return true
}
