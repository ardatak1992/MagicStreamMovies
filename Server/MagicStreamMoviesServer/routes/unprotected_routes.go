package routes

import (
	controller "github.com/ardatak1992/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUnprotectedRoutes(router *gin.Engine) {
	router.GET("/movies", controller.GetMovies())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
	
}
