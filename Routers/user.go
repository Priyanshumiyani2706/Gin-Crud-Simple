package Routers

import (
	Controller "gin-crud/Controllers"

	"github.com/gin-gonic/gin"
)

func User(router *gin.Engine) {

	// Get all user
	router.GET("/users", Controller.GetAllUsers)

	// Get one user
	router.GET("/user/:name", Controller.GetOneUser)

	// Update a user
	router.POST("/user/:name", Controller.UpdateUser)

	// Create a new user
	router.POST("/user", Controller.UserCreate)

	// Delete a user
	router.DELETE("/user/:name", Controller.DeleteUser)

}
