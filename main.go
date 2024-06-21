package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gbubemi22/gowith_ginsetup/src/app/database"
	"github.com/gbubemi22/gowith_ginsetup/src/app/middleware"
	"github.com/gbubemi22/gowith_ginsetup/src/app/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:" + err.Error())
	}

	  // Set Gin to release mode
	 // gin.SetMode(gin.ReleaseMode)

	// Initialize MongoDB client
	client, err := database.DBinstance()
	if err != nil {
		log.Fatalf("Error initializing MongoDB client: %v", err)
	}

	router := gin.New()
	router.Use(gin.Logger())


	routes.UserRouter(router)

	router.GET("/", func(c *gin.Context) {
		serviceName := os.Getenv("SERVICE_NAME")

		c.JSON(http.StatusOK, gin.H{
			"success":        true,
			"message":        "Welcome to My Service api",
			"httpStatusCode": http.StatusOK,
			"service":        serviceName,
		})
	})

	router.NoRoute(middleware.HandleNotFound)
	router.Use(middleware.ErrorHandlerMiddleware)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Use(func(c *gin.Context) {
		c.Set("mongoClient", client)
		c.Next()
	})

	router.Run(":" + port)
}
