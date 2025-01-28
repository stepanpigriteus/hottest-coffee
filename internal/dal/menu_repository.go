package dal

import (
	"encoding/json"
	"hot/internal/pkg/config"
	"hot/models"
	"io"
	"os"
	"path/filepath"
)

type MenuInterface interface {
	GetMenuItemById(id string) (models.MenuItem, error)
	DeleteMenuItemById(id string) error
	PutMenuItemById(id, name, description string, price float64, ingredients []models.MenuItemIngredient) error
	GetMenuItems() ([]models.MenuItem, error)
	PostMenuItem(menuItem *models.MenuItem) error
}

type MenuItems struct {
	menu []models.MenuItem
}

func (menuItems *MenuItems) GetMenuItemById(id string) (models.MenuItem, error) {
	if err := OpenMenu(menuItems); err != nil {
		return models.MenuItem{}, err
	}

	for _, el := range menuItems.menu {
		if el.ID == id {
			return el, nil
		}
	}

	return models.MenuItem{}, models.ErrItemNotFound
}

func (menuItems *MenuItems) GetMenuItems() ([]models.MenuItem, error) {
	if err := OpenMenu(menuItems); err != nil {
		return nil, err
	}

	return menuItems.menu, nil
}

func (menuItems *MenuItems) DeleteMenuItemById(id string) error {
	if err := OpenMenu(menuItems); err != nil {
		return err
	}

	var indexToDelete int = -1

	for i, el := range menuItems.menu {
		if el.ID == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete == -1 {
		return models.ErrItemNotFound
	}

	menuItems.menu = append(menuItems.menu[:indexToDelete], menuItems.menu[indexToDelete+1:]...)

	if err := saveMenuToFile(menuItems); err != nil {
		return err
	}

	return nil
}

func (menuItems *MenuItems) PostMenuItem(menuItem *models.MenuItem) error {
	if err := OpenMenu(menuItems); err != nil {
		return err
	}

	for _, item := range menuItems.menu {
		if item.ID == menuItem.ID {
			return models.ErrDuplicateMenuItemID
		}
	}

	menuItems.menu = append(menuItems.menu, *menuItem)

	if err := saveMenuToFile(menuItems); err != nil {
		return err
	}

	return nil
}

func (menuItems *MenuItems) PutMenuItemById(id, name, description string, price float64, ingredients []models.MenuItemIngredient) error {
	if err := OpenMenu(menuItems); err != nil {
		return err
	}

	for i, el := range menuItems.menu {
		if el.ID == id {

			menuItems.menu[i].Name = name
			menuItems.menu[i].Description = description
			menuItems.menu[i].Price = price
			menuItems.menu[i].Ingredients = ingredients

			if err := saveMenuToFile(menuItems); err != nil {
				return err
			}

			return nil
		}
	}

	return models.ErrItemNotFound
}

func OpenMenu(menuItems *MenuItems) error {
	path := filepath.Join(config.Dir, "menu.json")

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	value, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(value, &menuItems.menu)
	if err != nil {
		return err
	}

	return nil
}

func saveMenuToFile(menuItems *MenuItems) error {
	path := filepath.Join(config.Dir, "menu.json")

	// Преобразуем меню в JSON
	data, err := json.Marshal(menuItems.menu)
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
