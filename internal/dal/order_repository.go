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
	PutOrderById(items *models.Order, id string) error
	GetOrders() ([]models.Order, error)
	PostOrder(item *models.Order) error
	CloseOrder(id string) (models.Order, error)
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

	if err := saveOrdersToFile(orders); err != nil {
		return err
	}

	return nil
}

func (orders *Orders) PostOrder(order *models.Order) error {
	err := CheckOrder(order)
	if err != nil {
		return err
	}

	if err := Open(orders); err != nil {
		return err
	}

	order.ID = generateNewOrderID(orders.orders)
	order.Status = "open"

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

func (orders *Orders) PutOrderById(item *models.Order, id string) error {
	err := CheckOrder(item)

	isExist := false
	if err != nil {
		return err
	}

	if err := Open(orders); err != nil {
		return err
	}

	for i, order := range orders.orders {

		if order.ID == id {
			isExist = true
			for _, r := range item.Items {
				if r.Quantity < 0 || r.Quantity > 10000 {
					err := fmt.Errorf("Incuficient: %s", r.ProductID)
					return err
				}
			}

			orders.orders[i].ID = id
			orders.orders[i].CustomerName = item.CustomerName
			orders.orders[i].Items = item.Items
			orders.orders[i].Status = item.Status
			orders.orders[i].CreatedAt = item.CreatedAt

			if err := saveOrdersToFile(orders); err != nil {
				return err
			}

			return nil
		}

		if !isExist {
			return models.ErrOrderNotFound
		}
	}

	return models.ErrOrderNotFound
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

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.Size() == 0 {
		return nil
	}

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

	for _, order := range orders {
		matches := re.FindStringSubmatch(order.ID)
		if len(matches) > 1 {
			id, err := strconv.Atoi(matches[1])
			if err == nil && id > maxID {
				maxID = id
			}
		}
	}

	newID := maxID + 1
	return fmt.Sprintf("order%d", newID)
}

func GetOrders() (Orders, error) {
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
	var orders Orders
	err = json.Unmarshal(value, &orders.orders)
	if err != nil {
		return Orders{}, err
	}

	return orders, nil
}

func CheckOrder(order *models.Order) error {
	menu := new(MenuItems)
	menuItems, err := menu.GetMenuItems()
	if err != nil {
		return err
	}
	inventory := new(Items)
	invItems, err := inventory.GetItems()
	if err != nil {
		return err
	}
	if err := ValidateOrderIngredients(order, menuItems, invItems); err != nil {
		fmt.Println(err)
		return err
	}

	for _, item := range order.Items {
		if item.ProductID == "" || item.Quantity <= 0 {
			err := fmt.Errorf("Incorrect value in order: %s quantity %d", item.ProductID, item.Quantity)
			return err
		}
	}
	return nil
}

func ValidateOrderIngredients(order *models.Order, menuItems []models.MenuItem, inventory []models.InventoryItem) error {
	inventoryMap := make(map[string]int)
	for _, item := range inventory {
		inventoryMap[item.IngredientID] = int(item.Quantity)
	}

	menuMap := make(map[string]models.MenuItem)

	for _, item := range menuItems {
		menuMap[item.ID] = item
	}

	for _, orderItem := range order.Items {

		menuItem, exists := menuMap[orderItem.ProductID]

		if !exists {
			err := fmt.Errorf("product not found: %s", orderItem.ProductID)
			return err
		}

		for _, ingredient := range menuItem.Ingredients {
			requiredQty := int(ingredient.Quantity) * orderItem.Quantity
			currentQty, exists := inventoryMap[ingredient.IngredientID]
			if !exists || currentQty < requiredQty {
				return models.ErrInsufficientIngredients
			}

			inventoryMap[ingredient.IngredientID] -= requiredQty
		}
	}

	return nil
}
