package dal

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"hot/internal/pkg/config"
	"hot/models"
)

type ItemSale struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity_sold"`
	TotalSale float64 `json:"total_sale"`
}

func GetAggregations() (models.Aggregation, error) {
	pathAggr := filepath.Join(config.Dir, "aggregations.json")
	orders, err := GetOrders()
	count := CountClosedOrders(orders.orders)

	UpdateAggregations(count)
	jsonFileAggr, err := os.Open(pathAggr)
	if err != nil {
		return models.Aggregation{}, err
	}
	defer jsonFileAggr.Close()

	byteValueAggr, err := io.ReadAll(jsonFileAggr)
	if err != nil {
		return models.Aggregation{}, err
	}

	var aggreg models.Aggregation
	err = json.Unmarshal(byteValueAggr, &aggreg)
	if err != nil {
		return models.Aggregation{}, err
	}
	return aggreg, nil
}

func CountClosedOrders(orders []models.Order) float64 {
	var count float64 = 0
	menu := &MenuItems{}
	for _, order := range orders {
		if order.Status == "closed" {
			for _, r := range order.Items {
				item, _ := menu.GetMenuItemById(r.ProductID)
				count += item.Price * float64(r.Quantity)
			}
		}
	}
	return count
}

func UpdateAggregations(closedCount float64) error {
	pathAggr := filepath.Join(config.Dir, "aggregations.json")

	file, err := os.Open(pathAggr)
	if err != nil {
		return err
	}
	defer file.Close()

	var aggregation models.Aggregation
	if err := json.NewDecoder(file).Decode(&aggregation); err != nil {
		return err
	}

	aggregation.TotalSales = closedCount

	data, err := json.MarshalIndent(aggregation, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(pathAggr, data, 0o644)
}

func GetItemSales() ([]models.AggregationItem, error) {
	// Получаем заказы
	orders, err := GetOrders()
	if err != nil {
		return nil, err
	}

	// Создаем структуру для подсчета продаж каждого продукта
	menu := &MenuItems{}
	productSalesMap := make(map[string]*ItemSale)

	// Подсчитываем продажи по каждому продукту
	for _, order := range orders.orders {
		if order.Status == "closed" {
			for _, r := range order.Items {
				item, err := menu.GetMenuItemById(r.ProductID)
				if err != nil {
					continue // Пропускаем, если продукта нет в меню
				}

				// Обновляем продажи для данного продукта
				if _, exists := productSalesMap[r.ProductID]; exists {
					productSalesMap[r.ProductID].Quantity += r.Quantity
					productSalesMap[r.ProductID].TotalSale += item.Price * float64(r.Quantity)
				} else {
					productSalesMap[r.ProductID] = &ItemSale{
						ProductID: r.ProductID,
						Quantity:  r.Quantity,
						TotalSale: item.Price * float64(r.Quantity),
					}
				}
			}
		}
	}

	// Создаем массив для ItemSales
	itemSales := make([]models.AggregationItem, 0, len(productSalesMap))
	for _, sales := range productSalesMap {
		itemSales = append(itemSales, models.AggregationItem{
			ID:        sales.ProductID,
			QuantSold: sales.Quantity,
		})
	}

	// Путь к файлу с агрегациями
	pathAggr := filepath.Join(config.Dir, "aggregations.json")

	// Открываем файл агрегаций
	file, err := os.Open(pathAggr)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Декодируем существующие данные в структуру Aggregation
	var aggregation models.Aggregation
	if err := json.NewDecoder(file).Decode(&aggregation); err != nil {
		return nil, err
	}

	// Обновляем данные о продажах товаров в агрегации
	aggregation.ItemSales = itemSales

	// Перезаписываем файл с обновленной агрегацией
	data, err := json.MarshalIndent(aggregation, "", "\t")
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(pathAggr, data, 0o644); err != nil {
		return nil, err
	}

	// Возвращаем itemSales
	return itemSales, nil
}
