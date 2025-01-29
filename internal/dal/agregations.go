package dal

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"hot/internal/pkg/config"
	"hot/models"
)

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
