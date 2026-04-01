package modules

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	users := router.Group("/user")

	users.GET("", GetAllUsers)
	users.GET("/:id", GetUserByID)
	users.POST("", CreateUser)
	users.PUT("/:id", UpdateUser)
	users.DELETE("/:id", DeleteUser)
}