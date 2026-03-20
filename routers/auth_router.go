package routers

import (
	"todo-backend/controllers"
	"todo-backend/repositories"
	"todo-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(r *gin.Engine, db *gorm.DB) {
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	group := r.Group("api/v1/auth")
	group.POST("/signup", authController.Signup)
}
