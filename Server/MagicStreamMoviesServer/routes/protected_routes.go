package routes

import (
	controller "github.com/ardatak1992/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	"github.com/ardatak1992/MagicStreamMovies/Server/MagicStreamMoviesServer/middleware"
	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())

}
