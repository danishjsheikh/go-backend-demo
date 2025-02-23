package controllers

import (
	"net/http"

	"github.com/danishjsheikh/go-backend-demo/app/core"
	"github.com/danishjsheikh/go-backend-demo/app/models"

	"github.com/gin-gonic/gin"
)

type OrderController struct{}

// GetOrderByCode retrieves an order by its order code
// @Summary Get Order by Code
// @Description Retrieves an order by order code
// @Tags Orders
// @Accept json
// @Produce json
// @Param x-correlationid header string true "Correlation ID"
// @Param orderCode path string true "Order Code"
// @Success 200 {object} models.Order
// @Failure 404 {object} models.ErrResponse
// @Router /orders/{orderCode} [get]
func (oc *OrderController) GetOrderByCode(c *gin.Context) {
	order, err := core.GetOrder(c.Param("orderCode"))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrResponse{Message: "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// CreateOrder creates a new order
// @Summary Create Order
// @Description Creates a new order
// @Tags Orders
// @Accept json
// @Produce json
// @Param x-correlationid header string true "Correlation ID"
// @Param request body models.CreateOrderRequest true "Order Request"
// @Success 201 {object} models.Order
// @Failure 400 {object} models.ErrResponse
// @Router /orders [post]
func (oc *OrderController) CreateOrder(c *gin.Context) {
	var request models.CreateOrderRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Message: err.Error()})
		return
	}
	order, err := core.CreateOrder(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}
