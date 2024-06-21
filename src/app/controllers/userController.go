package controllers

import (
	"net/http"

	"github.com/gbubemi22/gowith_ginsetup/src/app/model"
	"github.com/gbubemi22/gowith_ginsetup/src/app/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	userService := service.NewUserService()
	return &UserController{
		userService: userService,
	}
}

// CreateUserHandler handles HTTP POST requests to create a new user.
func (controller *UserController) CreateUserHandler(c *gin.Context) {
	var userInput model.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Call CreateUser method from UserService
	createdUser, err := controller.userService.CreateUser(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
