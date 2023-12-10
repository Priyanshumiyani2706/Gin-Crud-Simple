package Controller

import (
	"gin-crud/Config"
	Model "gin-crud/Models"

	"github.com/gin-gonic/gin"
)

// Create user in database
func UserCreate(c *gin.Context) {
	var user Model.User
	c.BindJSON(&user) // Getting json data from api and compare data with model
	if err := Config.DB.Create(user).Error; err != nil {
		c.JSON(404, err.Error())
	} else {
		c.JSON(200, user)
	}
}

// Get all users from the database
func GetAllUsers(c *gin.Context) {
	var users []Model.User
	if err := Config.DB.Find(&users).Error; err != nil {
		c.JSON(404, err.Error())
	} else {
		c.JSON(200, users)
	}
}

// Get one user with name
func GetOneUser(c *gin.Context) {
	var user Model.User
	name := c.Params.ByName("name")
	if err := Config.DB.Where("username = ?", name).First(&user).Error; err != nil {
		c.JSON(404, err.Error())
	} else {
		c.JSON(200, user)
	}
}

// Update user with name
func UpdateUser(c *gin.Context) {
	var user Model.User
	c.BindJSON(&user)
	name := c.Params.ByName("name")
	if err := Config.DB.Where("username = ?", name).Updates(user).Error; err != nil {
		c.JSON(404, err.Error())
	} else {
		c.JSON(200, user)
	}
}

// Delete user with name
func DeleteUser(c *gin.Context) {
	var user Model.User
	c.BindJSON(&user)
	name := c.Params.ByName("name")
	if err := Config.DB.Where("username = ?", name).Delete(&user).Error; err != nil {
		c.JSON(404, err.Error())
	} else {
		c.JSON(200, user)
	}
}
