package dal

import (
	"encoding/json"
	"fmt"
	"hot/config"
	"hot/models"
	"io"
	"os"
	"path/filepath"
)

type OrderInterface interface {
	GetOrderById(id string) (models.Order, error)
	DeleteOrderById(id string) error
	PutOrderById(id, customerName, status, createdAt string, items []models.OrderItem) error
	GetOrders() ([]models.Order, error)
	PostOrder(item *models.Order) error
	CloseOrder(id string) error
}

type Orders struct {
	orders []models.Order
}

func (orders *Orders) GetOrderById(id string) (models.Order, error) {
	path := filepath.Join(config.Dir, "orders.json")

	file, err := os.Open(path)
	if err != nil {
		return models.Order{}, err
	}
	defer file.Close()

	value, err := io.ReadAll(file)
	if err != nil {
		return models.Order{}, err
	}

	// джесоним гада
	err = json.Unmarshal(value, &orders.orders)
	if err != nil {
		return models.Order{}, err
	}

	for _, el := range orders.orders {
		fmt.Println(el)
		if el.ID == id {
			item := &models.Order{
				ID:           el.ID,
				CustomerName: el.CustomerName,
				Items:        el.Items,
				Status:       el.Status,
				CreatedAt:    el.CreatedAt,
			}
			return *item, nil
		}
	}
	return models.Order{}, models.ErrOrderNotFound
}
