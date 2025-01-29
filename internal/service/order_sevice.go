package service

import (
	"hot/internal/dal"
	"hot/models"
	"regexp"
	"strings"
)

func CreateNewOrder(order models.Order) (int, interface{}) {
	if order.CustomerName == "" || len(order.Items) == 0 {
		return 400, models.Error{Message: "Not enough arguments"}
	}

	name := strings.Trim(order.CustomerName, " ")
	pattern := `^[A-Za-z]+ [A-Za-z]+$`
	re := regexp.MustCompile(pattern)

	if !re.MatchString(name) {
		return 404, models.Error{Message: "invalid name, write both name and last name using only letters"}
	}
	menu, err := dal.GetMenu()
	if err != nil {
		return 400, models.Error{Message: "internal server error"}
	}
	menu2, err2 := menu.GetMenuItems()
	if err2 != nil {
		return 400, models.Error{Message: "internal server error"}
	}
	for _, item := range order.Items {
		if !menuExist(menu2, item.ProductID) {
			return 400, models.Error{Message: "Unknown position " + item.ProductID}
		}
		if item.Quantity < 1 {
			return 400, models.Error{Message: "Invalid quantity for position"}
		}
	}
	inventory, errI := dal.GetInventory()
	if errI != nil {
		return 400, models.Error{Message: "internal server error"}
	}
	inventory2, errI2 := inventory.GetItems()
	if errI2 != nil {
		return 400, models.Error{Message: "internal server error"}
	}
	orderNeeds := make(map[string]float64)
	for _, position := range order.Items {
	}
}

func menuExist(menu []models.MenuItem, product string) bool {
	for _, item := range menu {
		if item.Name == product {
			return true
		}
	}
	return false
}
