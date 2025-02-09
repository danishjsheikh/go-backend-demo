package models

// UserResponse
// @Description User Valid Response
type UserResponse struct {
	Data []User `json:"data"`
}

// ErrResponse
// @Description User Error Response
type ErrResponse struct {
	Message string
}

// User
// @Description User Schema
type User struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

// CreateOrderRequest
// @Description Request about creating Order
type CreateOrderRequest struct {
	// shipment no of Order
	ShipmentNumber string `json:"shipmentNumber" validate:"required"`
	// country code like: tr, us
	CountryCode string `json:"countryCode" validate:"required,len=2"`
	// age to make sure you are young
	Age int `json:"age" validate:"required,oldAge"`
}
