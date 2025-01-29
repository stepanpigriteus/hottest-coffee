package dal

import (
	"encoding/json"
	"fmt"
	"hot/internal/pkg/config"
	"hot/models"
	"io"
	"os"
	"path/filepath"
)

type ItemInterface interface {
	GetItemById(id string) (models.InventoryItem, error)
	DeleteItemById(id string) error
	PutItemById(id, name, unit string, quantity float64) error
	GetItems() ([]models.InventoryItem, error)
	PostItem(item *models.InventoryItem) error
}

type Items struct {
	inventory []models.InventoryItem
}

func (items *Items) GetItemById(id string) (models.InventoryItem, error) {
	// Открываем файл с инвентарем
	if err := OpenItems(items); err != nil {
		return models.InventoryItem{}, err
	}

	// Ищем товар по ID
	for _, item := range items.inventory {
		if item.IngredientID == id {
			return item, nil
		}
	}

	return models.InventoryItem{}, models.ErrItemNotFound
}

func (items *Items) GetItems() ([]models.InventoryItem, error) {
	if err := OpenItems(items); err != nil {
		return nil, err
	}

	return items.inventory, nil
}

func (items *Items) DeleteItemById(id string) error {
	if err := OpenItems(items); err != nil {
		return err
	}

	var indexToDelete int = -1
	for i, item := range items.inventory {
		if item.IngredientID == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete == -1 {
		return models.ErrItemNotFound
	}

	items.inventory = append(items.inventory[:indexToDelete], items.inventory[indexToDelete+1:]...)

	if err := saveItemsToFile(items); err != nil {
		return err
	}

	return nil
}

func (items *Items) PostItem(item *models.InventoryItem) error {
	if err := OpenItems(items); err != nil {
		return err
	}

	for _, existingItem := range items.inventory {
		if existingItem.IngredientID == item.IngredientID {
			return models.ErrDuplicateInventoryItemID
		}
	}

	items.inventory = append(items.inventory, *item)
	if err := saveItemsToFile(items); err != nil {
		return err
	}
	return nil
}

func (items *Items) PutItemById(o models.InventoryItem, id string) error {
	if err := OpenItems(items); err != nil {
		return err
	}

	for i, item := range items.inventory {
		if item.IngredientID == id {

			items.inventory[i].Name = o.Name
			items.inventory[i].Unit = o.Unit
			items.inventory[i].Quantity = o.Quantity
			fmt.Println(items.inventory[i], item.Quantity)
			if err := saveItemsToFile(items); err != nil {
				return err
			}

			return nil
		}
	}

	return models.ErrItemNotFound
}

func OpenItems(items *Items) error {
	path := filepath.Join(config.Dir, "inventory.json")

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	value, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(value, &items.inventory)
	if err != nil {
		return err
	}

	return nil
}

func saveItemsToFile(items *Items) error {
	path := filepath.Join(config.Dir, "inventory.json")

	data, err := json.Marshal(items.inventory)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func GetInventory() (Items, error) {
	path := filepath.Join(config.Dir, "inventory.json")

	file, err := os.Open(path)
	if err != nil {
		return Items{}, err
	}
	defer file.Close()

	value, err := io.ReadAll(file)
	if err != nil {
		return Items{}, err
	}
	var items Items
	err = json.Unmarshal(value, &items.inventory)

	if err != nil {
		return Items{}, err
	}

	return items, nil
}
