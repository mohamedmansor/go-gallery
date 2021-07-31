package main

import (
	"blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		models.LoadEnv()
		models.NewDatabase()
		context.JSON(http.StatusOK, gin.H{"data": "Hello"})
	})
	router.Run(":8000")
}
