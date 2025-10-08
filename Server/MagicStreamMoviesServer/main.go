package main

import (
	"fmt"

	controller "github.com/ardatak1992/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies")
	})

	router.GET("/movies", controller.GetMovies())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}

}
