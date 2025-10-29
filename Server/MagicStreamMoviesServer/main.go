package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ardatak1992/MagicStreamMovies/Server/MagicStreamMoviesServer/database"
	"github.com/ardatak1992/MagicStreamMovies/Server/MagicStreamMoviesServer/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies")
	})

	var client *mongo.Client = database.Connect()

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to find .env file")
	}

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	var origins []string
	if allowedOrigins != "" {
		origins = strings.Split(allowedOrigins, ",")
		for i := range origins {
			origins[i] = strings.TrimSpace(origins[i])
			log.Println("Allowed Origin:", origins[i])
		}
	} else {
		origins = []string{"http://localhost:5173"}
		log.Println("Allowed Origin:", origins[0])
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to reach server: %v", err)
	}

	defer func() {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Fatalf("Failed to disconnect form MongoDB: %v", err)
		}
	}()

	config := cors.Config{}
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PATCH"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))
	router.Use(gin.Logger())

	routes.SetupUnprotectedRoutes(router, client)
	routes.SetupProtectedRoutes(router, client)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}

}
