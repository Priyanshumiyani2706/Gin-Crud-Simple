package main

import (
	"gin-crud/Config"
	"gin-crud/Routers"

	"github.com/gin-gonic/gin"
)

func init() {
	Config.LoadEnv()
	// Getting port 3000 form env file and for slect other things os.Getenv(key)
	Config.DbConnection()
	// Database connection
}

func main() {
	r := gin.Default() //Create server

	Routers.User(r) // Pass routing in user routes

	r.Run() // run server
}
