package dal

import (
	"encoding/json"
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
	// Открываем файл с инвентарем
	if err := OpenItems(items); err != nil {
		return nil, err
	}

	// Возвращаем все товары
	return items.inventory, nil
}

func (items *Items) DeleteItemById(id string) error {
	// Открываем файл с инвентарем
	if err := OpenItems(items); err != nil {
		return err
	}

	// Ищем индекс товара для удаления
	var indexToDelete int = -1
	for i, item := range items.inventory {
		if item.IngredientID == id {
			indexToDelete = i
			break
		}
	}

	// Если товар не найден
	if indexToDelete == -1 {
		return models.ErrItemNotFound
	}

	// Удаляем товар
	items.inventory = append(items.inventory[:indexToDelete], items.inventory[indexToDelete+1:]...)

	// Сохраняем обновленный инвентарь в файл
	if err := saveItemsToFile(items); err != nil {
		return err
	}

	return nil
}

func (items *Items) PostItem(item *models.InventoryItem) error {
	// Открываем файл с инвентарем
	if err := OpenItems(items); err != nil {
		return err
	}

	// Проверяем, есть ли товар с таким ID
	for _, existingItem := range items.inventory {
		if existingItem.IngredientID == item.IngredientID {
			return models.ErrDuplicateInventoryItemID
		}
	}

	// Добавляем новый товар
	items.inventory = append(items.inventory, *item)

	// Сохраняем обновленный инвентарь в файл
	if err := saveItemsToFile(items); err != nil {
		return err
	}

	return nil
}

func (items *Items) PutItemById(id, name, unit string, quantity float64) error {
	// Открываем файл с инвентарем
	if err := OpenItems(items); err != nil {
		return err
	}

	// Ищем товар по ID
	for i, item := range items.inventory {
		if item.IngredientID == id {
			// Обновляем данные товара
			items.inventory[i].Name = name
			items.inventory[i].Unit = unit
			items.inventory[i].Quantity = quantity

			// Сохраняем обновленный инвентарь в файл
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

	// Открываем файл с инвентарем
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Читаем все данные из файла
	value, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Десериализуем данные в inventory
	err = json.Unmarshal(value, &items.inventory)
	if err != nil {
		return err
	}

	return nil
}

func saveItemsToFile(items *Items) error {
	path := filepath.Join(config.Dir, "inventory.json")

	// Преобразуем инвентарь в JSON
	data, err := json.Marshal(items.inventory)
	if err != nil {
		return err
	}

	// Открываем файл для записи
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Записываем данные в файл
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
