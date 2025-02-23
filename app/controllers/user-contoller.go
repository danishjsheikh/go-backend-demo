package controllers

import (
	"log"
	"net/http"

	"github.com/danishjsheikh/go-backend-demo/app/core"
	"github.com/danishjsheikh/go-backend-demo/app/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// GetUsers return list of all users from the database
// @Summary return list of all users
// @Description return list of all users from the database
// @Tags Users
// @Success 200 {object} models.UserResponse
// @Failure 500 {object} models.ErrResponse
// @Router /users [get]
func (uc *UserController) GetUsers(c *gin.Context) {

	users, err := core.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{Message: err.Error()})
		return
	}

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
// @Failure 400 {object} models.ErrResponse
// @Router /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {

	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Message: err.Error()})
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

// Additional CRUD methods for Users
// @Summary Update User
// @Description Updates an existing user
// @Tags Users
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Param request body models.User true "Updated User"
// @Success 200 {object} models.User
// @Failure 404 {object} models.ErrResponse
// @Router /users/{username} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	var request models.User
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Message: err.Error()})
		return
	}
	user, err := core.UpdateUser(c.Param("username"), request)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Delete User
// @Description Deletes a user
// @Tags Users
// @Param username path string true "Username"
// @Success 204 {string} string "No Content"
// @Failure 404 {object} models.ErrResponse
// @Router /users/{username} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	err := core.DeleteUser(c.Param("username"))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, "No Content")
}
