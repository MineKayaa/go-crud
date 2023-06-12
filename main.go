package main

import (
	"github.com/MineKayaa/go-crud/controllers"
	"github.com/MineKayaa/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {

	r := gin.Default()

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run()
}
