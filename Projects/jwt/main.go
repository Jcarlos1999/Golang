package main

import (
	"github.com/Jcarlos1999/Golang/Projects/jwt/controllers"
	"github.com/Jcarlos1999/Golang/Projects/jwt/initializers"
	"github.com/Jcarlos1999/Golang/Projects/jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
}

// / testes
func main() {
	r := gin.Default()
	/// testes
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validade", middleware.RequireAuth, controllers.Validade)
	r.Run()
}
