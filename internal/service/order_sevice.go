package service

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"hot/internal/dal"
	"hot/models"
)

func GetOrderById(id string) (models.Order, int, string) {
	var orderRepo dal.OrderInterface = new(dal.Orders)
	order, err := orderRepo.GetOrderById(id)
	if err != nil {
		if err == models.ErrOrderNotFound {
			return models.Order{}, http.StatusNotFound, err.Error()
		}
		return models.Order{}, http.StatusInternalServerError, ""
	}

	return order, http.StatusOK, ""
}

func DeleteOrderById(id string) (int, string) {
	var orderRepo dal.OrderInterface = new(dal.Orders)
	err := orderRepo.DeleteOrderById(id)
	if err != nil {
		if err == models.ErrMenuItemNotFound {
			return http.StatusNotFound, err.Error()
		}
		return http.StatusInternalServerError, ""
	}

	return http.StatusOK, ""
}

func PutOrderById(order *models.Order, id string) (int, string) {
	if !ValidOrder(order) {
		return http.StatusBadRequest, models.ErrInvalidOrderItem.Error()
	}

	var orderRepo dal.OrderInterface = new(dal.Orders)

	err := orderRepo.PutOrderById(order, id)
	if err != nil {
		if err == models.ErrOrderNotFound {
			return http.StatusNotFound, err.Error()
		} else if err == models.ErrOrderClosed {
			return http.StatusBadRequest, err.Error()
		}
		return http.StatusInternalServerError, ""
	}

	return http.StatusOK, ""
}

func GetOrders() ([]models.Order, int) {
	var orderRepo dal.OrderInterface = new(dal.Orders)
	orders, err := orderRepo.GetOrders()
	if err != nil {
		return make([]models.Order, 0), http.StatusInternalServerError
	}

	return orders, http.StatusOK
}

func PostOrder(item *models.Order) (int, string) {
	if !ValidOrder(item) {
		return http.StatusBadRequest, models.ErrInvalidOrderItem.Error()
	}

	item.CreatedAt = time.Now().Format(time.RFC3339)

	var orderRepo dal.OrderInterface = new(dal.Orders)
	err := orderRepo.PostOrder(item)
	if err != nil {
		switch err {
		case models.ErrDuplicateOrderID:
			item.ID = strconv.Itoa(rand.Intn(10000000))
			return PostOrder(item)
		case models.ErrMenuItemNotFound:
			return http.StatusNotFound, err.Error()
		case models.ErrInsufficientIngredients:
			return http.StatusConflict, err.Error()
		default:
			return http.StatusInternalServerError, err.Error()
		}
	}

	return http.StatusCreated, "Order created successfully"
}

func CloseOrder(id string) (int, string) {
	var orderRepo dal.OrderInterface = new(dal.Orders)
	order, err := orderRepo.CloseOrder(id)
	if err != nil {

		if err == models.ErrOrderNotFound {
			return http.StatusNotFound, err.Error()
		} else if err == models.ErrOrderClosed {
			return http.StatusBadRequest, err.Error()
		} else if err == models.ErrInsufficientIngredients {
			return http.StatusConflict, err.Error()
		}
		return http.StatusInternalServerError, ""
	}
	if order.Status == "closed" {
		return http.StatusOK, "Order successfully closed"
	}
	return http.StatusOK, ""
}

func ValidOrder(order *models.Order) bool {
	if order == nil {
		return false
	}

	if order.CustomerName == "" ||
		len(order.Items) == 0 {
		return false
	}

	return true
}
