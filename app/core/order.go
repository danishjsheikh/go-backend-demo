package core

import (
	"errors"
	"sync"

	"github.com/danishjsheikh/go-backend-demo/app/models"
)

var (
	orders = make(map[string]models.Order)
	mutex  sync.Mutex
)

// GetOrder retrieves an order by its order code
func GetOrder(orderCode string) (*models.Order, error) {
	mutex.Lock()
	defer mutex.Unlock()

	order, exists := orders[orderCode]
	if !exists {
		return nil, errors.New("order not found")
	}
	return &order, nil
}

// CreateOrder creates a new order and stores it
func CreateOrder(request models.CreateOrderRequest) (*models.Order, error) {
	mutex.Lock()
	defer mutex.Unlock()

	order := models.Order{
		ShipmentNumber: request.ShipmentNumber,
		CountryCode:    request.CountryCode,
		Age:            request.Age,
	}
	orders[request.ShipmentNumber] = order
	return &order, nil
}

// UpdateOrder updates an existing order
func UpdateOrder(orderCode string, request models.CreateOrderRequest) (*models.Order, error) {
	mutex.Lock()
	defer mutex.Unlock()

	_, exists := orders[orderCode]
	if !exists {
		return nil, errors.New("order not found")
	}

	updatedOrder := models.Order{
		ShipmentNumber: request.ShipmentNumber,
		CountryCode:    request.CountryCode,
		Age:            request.Age,
	}
	orders[orderCode] = updatedOrder
	return &updatedOrder, nil
}

// DeleteOrder removes an order
func DeleteOrder(orderCode string) error {
	mutex.Lock()
	defer mutex.Unlock()

	_, exists := orders[orderCode]
	if !exists {
		return errors.New("order not found")
	}
	delete(orders, orderCode)
	return nil
}
