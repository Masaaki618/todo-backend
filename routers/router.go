package routers

import (
	"todo-backend/infra"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	db := infra.SetupDB()
	r := gin.Default()

	SetupItemRoutes(r, db)
	SetupAuthRoutes(r, db)

	return r
}
