package Routers

import (
	"fmt"
	Controller "gin-crud/Controllers"

	"github.com/gin-gonic/gin"
)

func User(router *gin.Engine) {
	fmt.Println("out of func")
	router.GET("/users", Controller.GetAllUsers)
	router.GET("/user/:name", Controller.GetOneUser)

	router.POST("/user/:name", Controller.UpdateUser)

	router.POST("/user", Controller.UserCreate)

	router.DELETE("/user/:name", Controller.DeleteUser)

}
