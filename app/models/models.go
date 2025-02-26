package models

// User represents a user schema
// @Description User Schema
type User struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

// UserResponse represents a valid response for user operations
// @Description User Valid Response
type UserResponse struct {
	Data []User `json:"data"`
}

// ErrResponse represents an error response
// @Description User Error Response
type ErrResponse struct {
	Message string `json:"error"`
}

// CreateOrderRequest represents a request for creating an order
// @Description Request about creating Order
type CreateOrderRequest struct {
	OrderID string `json:"order_id" validate:"required"`
	// shipment no of Order
	ShipmentNumber string `json:"shipmentNumber" validate:"required"`
	// country code like: tr, us
	CountryCode string `json:"countryCode" validate:"required,len=2"`
	// age to make sure you are young
	Age int `json:"age" validate:"required,oldAge"`
}

// Order represents an order schema
// @Description Order Schema
type Order struct {
	OrderID        string `json:"order_id"`
	ShipmentNumber string `json:"shipmentNumber"`
	CountryCode    string `json:"countryCode"`
	Age            int    `json:"age"`
}
