package main

import (
	"log"
	"os"
	"products/config"
	_ "products/docs" // Import the docs package for Swagger
	"products/handlers"
	"products/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go API with JWT and MongoDB
// @version         1.0
// @description     This is a simple API that demonstrates JWT authentication and MongoDB integration.
// @host            localhost:8080
// @BasePath        /
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Access environment variables
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")
	port := os.Getenv("SERVER_PORT")

	r := gin.Default()
	db := config.ConnectDB(mongoURI, dbName)

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/login", handlers.LoginHandler(db, jwtSecret))

	r.POST("/register", handlers.RegisterHandler(db))

	authorized := r.Group("/")
	authorized.Use(middlewares.JWTAuthMiddleware(jwtSecret))
	{
		authorized.POST("/items", handlers.CreateItemHandler(db))
		authorized.GET("/items", handlers.GetItemsHandler(db))
		authorized.GET("/items/:id", handlers.GetItemDetailHandler(db))
	}

	r.Run(":" + port)
}
