package dal

import (
	"encoding/json"
	"hot/internal/pkg/config"
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
	Open(orders)
	for _, el := range orders.orders {
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

func (orders *Orders) DeleteOrderById(id string) error {
	if _, err := Open(orders); err != nil {
		return err
	}

	var indexToDelete int = -1
	for i, el := range orders.orders {
		if el.ID == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete == -1 {
		return models.ErrOrderNotFound
	}

	orders.orders = append(orders.orders[:indexToDelete], orders.orders[indexToDelete+1:]...)

	if err := saveOrdersToFile(orders); err != nil {
		return err
	}

	return nil
}

func (orders *Orders) UpdateOrderById(id string) error {
	return models.ErrDuplicateOrderID
}

func (orders *Orders) CloseOrderById(id string) error {
	return models.ErrDuplicateOrderID
}

func Open(orders *Orders) (Orders, error) {
	path := filepath.Join(config.Dir, "orders.json")

	file, err := os.Open(path)
	if err != nil {
		return Orders{}, err
	}
	defer file.Close()

	value, err := io.ReadAll(file)
	if err != nil {
		return Orders{}, err
	}

	err = json.Unmarshal(value, &orders)
	if err != nil {
		return Orders{}, err
	}

	return *orders, nil
}

func saveOrdersToFile(orders *Orders) error {
	path := filepath.Join(config.Dir, "orders.json")

	// Преобразуем orders в JSON
	data, err := json.Marshal(orders)
	if err != nil {
		return err
	}

	// Открываем файл для записи (перезаписываем его)
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
