package models

import "errors"

var (
	ErrInventoryFileNotFound    = errors.New("inventory file not found")
	ErrItemNotFound             = errors.New("item not found")
	ErrDuplicateInventoryItemID = errors.New("item with the same IngredientID already exists")
	ErrInvalidIntentoryItem     = errors.New("invalid inventory item")


	ErrMenuFileNotFound    = errors.New("menu items file not found")
	ErrMenuItemNotFound    = errors.New("menu item not found")
	ErrDuplicateMenuItemID = errors.New("menu item with this ID already exists")
	ErrInvalidMenuItem     = errors.New("invalid menu item")

	ErrOrderFileNotFound       = errors.New("orders file not found")
	ErrOrderNotFound           = errors.New("order not found")
	ErrDuplicateOrderID        = errors.New("order with this ID already exists")
	ErrInsufficientIngredients = errors.New("not enough ingredients for the order")
	ErrInvalidOrderItem        = errors.New("invalid order item")
	ErrOrderClosed             = errors.New("order is already closed")
	ErrOrderDelete             = errors.New("error order delete")
	
)


