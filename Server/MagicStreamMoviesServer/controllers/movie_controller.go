package controllers

import "github.com/gin-gonic/gin"

func GetMovies() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "List of Movies"})
	}
}
