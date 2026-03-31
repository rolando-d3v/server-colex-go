package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"

	. "server-colex-go/config"
	. "server-colex-go/modules/user"
)

func main() {
	// 🔥 Cargar .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No se encontró .env, usando variables del sistema")
	}

	// Conectar a la base de datos
	if err := InitDB(); err != nil {
		log.Fatalf("Error conectando a BD: %v", err)
	}

	// Configurar Gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Archivos estaticos
	router.Static("/public", "./public")

	// Middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Next()
	})

	// Rutas
	UserRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4002"
	}

	// Graceful shutdown
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		log.Println("🛑 Cerrando servidor...")
		CloseDB()
		os.Exit(0)
	}()

	log.Printf("🚀 Servidor ejecutando en puerto %s\n", port)
	router.Run(":" + port)
}
