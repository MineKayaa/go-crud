package main

import (
	"github.com/MineKayaa/go-crud/initializers"
	"github.com/MineKayaa/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
