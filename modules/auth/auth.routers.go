package modules

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	users := router.Group("/auth")

	users.POST("/login", AuthLogin)
	users.GET("/verify-auth", VerifyAuth)
	
	// mock de logout para borrar la cookie
	users.POST("/logout", func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "", false, true)
		c.JSON(200, gin.H{"msj": "Logout exitoso"})
	})
}