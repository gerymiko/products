package main

import (
	"products/config"
	_ "products/docs" // Import the docs package for Swagger
	"products/handlers"
	"products/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go API with JWT and MongoDB
// @version         1.0
// @description     This is a simple API that demonstrates JWT authentication and MongoDB integration.
// @host            localhost:8080
// @BasePath        /
func main() {
	r := gin.Default()
	db := config.ConnectDB()

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/login", handlers.LoginHandler(db))
	r.POST("/register", handlers.RegisterHandler(db))

	authorized := r.Group("/")
	authorized.Use(middlewares.JWTAuthMiddleware())
	{
		authorized.GET("/items", handlers.GetItemsHandler(db))
		authorized.GET("/items/:id", handlers.GetItemDetailHandler(db))
	}

	r.Run(":8080")
}
