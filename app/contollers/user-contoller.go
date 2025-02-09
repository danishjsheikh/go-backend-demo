package controllers

import (
	"backend-go/app/core"
	"backend-go/app/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// GetUsers return list of all users from the database
// @Summary return list of all
// @Description return list of all users from the database
// @Tags Users
// @Success 200 {object} models.UserResponse
// @Router /users [get]
func (uc *UserController) GetUsers(c *gin.Context) {

	users := core.ListUsers()

	c.JSON(http.StatusOK, models.UserResponse{Data: users})
}

// CreateUser create new user
// @Summary create new user
// @Description create new user
// @Security bearerToken
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 201 {object} models.User
// @Failure 400 {object}  models.ErrResponse
// @Router /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {

	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := core.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Message: err.Error()})
		return
	}

	auth := c.GetHeader("Authorization")
	log.Println("Authorization", auth)

	c.JSON(http.StatusCreated, req)
}
