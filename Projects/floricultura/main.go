package main

import (
	"gihub.com/Jcarlos1999/Golang/Projects/floricultura/internal/controllers"
	"gihub.com/Jcarlos1999/Golang/Projects/floricultura/internal/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvEnvironment()
	initializers.Conn()
	initializers.Sync()
}

func main() {

	// flower, _ := models.NewFlower("123", "Orchid", 11, time.Now())
	// flower.ToString()
	server := gin.Default()
	server.POST("/flowers/create", controllers.CreateFlower)
	server.POST("/flowers/find", controllers.FindFlowers)
	server.POST("/order/create", controllers.CreateOrder)
	server.Run()

}
