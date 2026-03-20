package controllers

import (
	"net/http"
	"todo-backend/dto"
	"todo-backend/services"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Signup(c *gin.Context)
}

type AuthController struct {
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{service: service}
}

func (a *AuthController) Signup(c *gin.Context) {
	var input dto.SignUpInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := a.service.SignUp(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.Status(http.StatusCreated)
}
