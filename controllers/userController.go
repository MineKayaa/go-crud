package controllers

import (
	"github.com/MineKayaa/go-crud/initializers"
	"github.com/MineKayaa/go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context) {

	//get id from url
	id := c.Param("id")

	var user models.User
	initializers.DB.First(&user, id)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func CreateUser(c *gin.Context) {
	//get data
	var body struct {
		Name string
		Age  int
	}

	c.Bind(&body)

	//create a user
	user := models.User{Name: body.Name, Age: body.Age}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UpdateUser(c *gin.Context) {

	//get id from url
	id := c.Param("id")

	//get data
	var body struct {
		Name string
		Age  int
	}

	c.Bind(&body)

	//find the user
	var user models.User
	initializers.DB.First(&user, id)

	//update
	initializers.DB.Model(&user).Updates(models.User{
		Name: body.Name,
		Age:  body.Age,
	})

	c.JSON(200, gin.H{
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {

	//get id from url
	id := c.Param("id")

	var user models.User
	initializers.DB.Delete(&user, id)

	c.Status(200)
}
