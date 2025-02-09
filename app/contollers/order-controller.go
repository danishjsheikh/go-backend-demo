package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danishjsheikh/go-backend-demo/app/core"
	"github.com/danishjsheikh/go-backend-demo/app/models"

	"github.com/gin-gonic/gin"
)

type OrderController struct{}

// GetOrderByCode Getting Order by Code
//
//	@Summary		Getting Order by Code
//	@Description	Getting Order by Code in detail
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			x-correlationid	header		string	true	"code of Order"
//	@Param			orderCode		path		string	true	"code of Order"
//	@Success		200				{string}	string
//	@Router			/orders/{orderCode} [get]
func (uc *OrderController) GetOrderByCode(c *gin.Context) {

	c.String(http.StatusOK, core.GetOrderByCode(fmt.Sprint(c.MustGet("correlationId")), c.Param("orderCode")))
}

// CreateOrder Creating Order
//
//	@Summary		Creating Order
//	@Description	Creating Order with given request
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			x-correlationid	header		string						true	"code of Order"
//	@Param			request			body		models.CreateOrderRequest	true	"Request of Creating Order Object"
//	@Success		200				{string}	string
//	@Failure		400				{string}	string	"Bad Request"
//	@Router			/orders [post]
func (uc *OrderController) CreateOrder(c *gin.Context) {
	var request models.CreateOrderRequest
	err := c.BindJSON(&request)
	if err != nil {
		log.Fatal("error binfing json", err)
	}
	message, err := core.CreateOrder(request)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.String(http.StatusOK, message)
	}

}
