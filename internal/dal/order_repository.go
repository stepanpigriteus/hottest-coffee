package dal

import (
	"encoding/json"
	"fmt"
	"hot/internal/pkg/config"
	"hot/models"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
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
	fmt.Println(id)
	if err := Open(orders); err != nil {
		fmt.Println(err)
		return models.Order{}, err
	}

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

func (orders *Orders) GetOrders() ([]models.Order, error) {
	if err := Open(orders); err != nil {
		return nil, err
	}

	return orders.orders, nil
}

func (orders *Orders) DeleteOrderById(id string) error {
	if err := Open(orders); err != nil {
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
	fmt.Println(orders.orders)
	if err := saveOrdersToFile(orders); err != nil {
		return err
	}

	return nil
}

func (orders *Orders) PostOrder(order *models.Order) error {
	if err := Open(orders); err != nil {
		return err
	}

	order.ID = generateNewOrderID(orders.orders)
	fmt.Println(order.ID)
	for _, existingOrder := range orders.orders {
		if existingOrder.ID == order.ID {
			return models.ErrDuplicateOrderID
		}
	}

	orders.orders = append(orders.orders, *order)

	if err := saveOrdersToFile(orders); err != nil {
		return err
	}

	return nil
}

// переделать
func (orders *Orders) PostUpdate(item *models.Order) error {
	if err := Open(orders); err != nil {
		return err
	}

	for _, order := range orders.orders {
		if order.ID == item.ID {
			return models.ErrDuplicateOrderID
		}
	}
	orders.orders = append(orders.orders, *item)

	if err := saveOrdersToFile(orders); err != nil {
		return err
	}

	return nil
}

func (orders *Orders) CloseOrder(id string) (models.Order, error) {
	if err := Open(orders); err != nil {
		return models.Order{}, err
	}

	for i, el := range orders.orders {
		if el.ID == id {

			orders.orders[i].Status = "closed"

			if err := saveOrdersToFile(orders); err != nil {
				return models.Order{}, err
			}

			return orders.orders[i], nil
		}
	}

	return models.Order{}, models.ErrOrderNotFound
}

func Open(orders *Orders) error {
	path := filepath.Join(config.Dir, "orders.json")

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	value, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(value) == 0 {
		orders.orders = []models.Order{}
		return nil
	}
	err = json.Unmarshal(value, &orders.orders)
	if err != nil {
		return err
	}

	return nil
}

func saveOrdersToFile(orders *Orders) error {
	path := filepath.Join(config.Dir, "orders.json")

	data, err := json.Marshal(orders.orders)
	if err != nil {
		return err
	}

	fmt.Printf("Serialized data: %s\n", string(data))

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

func generateNewOrderID(orders []models.Order) string {
	var maxID int
	re := regexp.MustCompile(`order(\d+)`)

	if len(orders) == 0 {
		return "order1"
	}
	// Ищем максимальный номер заказа
	for _, order := range orders {
		matches := re.FindStringSubmatch(order.ID)
		if len(matches) > 1 {
			id, err := strconv.Atoi(matches[1])
			if err == nil && id > maxID {
				maxID = id
			}
		}
	}

	// Генерируем новый ID с увеличенным числовым значением
	newID := maxID + 1
	return fmt.Sprintf("order%d", newID)
}
