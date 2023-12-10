package Controller

import (
	"gin-crud/Config"
	Model "gin-crud/Models"

	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	var user Model.User
	c.BindJSON(&user)

	if err := Config.DB.Create(user).Error; err != nil {
		c.JSON(404, err.Error())
	} else {
		c.JSON(200, user)
	}
}

func GetAllUsers(c *gin.Context) {
	var users []Model.User
	if err := Config.DB.Find(&users).Error; err != nil {
		c.JSON(404, err.Error())
	} else {
		c.JSON(200, users)
	}
}

func GetOneUser(c *gin.Context) {
	var user Model.User
	name := c.Params.ByName("name")
	if err := Config.DB.Where("username = ?", name).First(&user).Error; err != nil {
		c.JSON(404, err.Error())
	} else {
		c.JSON(200, user)
	}
}

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
