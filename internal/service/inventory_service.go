package service

import (
	"net/http"

	"hot/internal/dal"
	"hot/models"
)

func GetItemById(id string) (models.InventoryItem, int, string) {
	var inventoryRepo dal.ItemInterface = new(dal.Items)
	item, err := inventoryRepo.GetItemById(id)
	if err != nil {
		if err == models.ErrItemNotFound {
			return models.InventoryItem{}, http.StatusNotFound, err.Error()
		}
		return models.InventoryItem{}, http.StatusInternalServerError, ""
	}

	return item, http.StatusOK, ""
}

func DeleteItemById(id string) (int, string) {
	var inventoryRepo dal.ItemInterface = new(dal.Items)
	err := inventoryRepo.DeleteItemById(id)
	if err != nil {
		if err == models.ErrItemNotFound {
			return http.StatusNotFound, err.Error()
		}
		return http.StatusInternalServerError, ""
	}

	return http.StatusOK, ""
}

func PutItemById(item *models.InventoryItem) (int, string) {
	if !isValidInventoryItem(item) {
		return http.StatusBadRequest, models.ErrInvalidIntentoryItem.Error()
	}

	var inventoryRepo dal.ItemInterface = new(dal.Items)
	err := inventoryRepo.PutItemById(*item, item.IngredientID)
	if err != nil {
		if err == models.ErrItemNotFound {
			return http.StatusNotFound, err.Error()
		}
		return http.StatusInternalServerError, ""
	}

	return http.StatusOK, ""
}

func GetItems() ([]models.InventoryItem, int) {
	var inventoryRepo dal.ItemInterface = new(dal.Items)
	items, err := inventoryRepo.GetItems()
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	return items, http.StatusOK
}

func PostItem(item *models.InventoryItem) (int, string) {
	if !isValidInventoryItem(item) {
		return http.StatusBadRequest, models.ErrInvalidIntentoryItem.Error()
	}

	var inventoryRepo dal.ItemInterface = new(dal.Items)
	err := inventoryRepo.PostItem(item)
	if err != nil {
		if err == models.ErrDuplicateInventoryItemID {
			return http.StatusBadRequest, err.Error()
		}
		return http.StatusInternalServerError, ""
	}

	return http.StatusCreated, ""
}

func isValidInventoryItem(inventoryItem *models.InventoryItem) bool {
	if inventoryItem == nil {
		return false
	}

	if inventoryItem.IngredientID == "" ||
		inventoryItem.Name == "" ||
		inventoryItem.Quantity <= 0 ||
		inventoryItem.Unit == "" {
		return false
	}

	return true
}
