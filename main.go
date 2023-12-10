package main

import (
	"gin-crud/Config"
	"gin-crud/Routers"

	"github.com/gin-gonic/gin"
)

func init() {
	Config.LoadEnv()
	Config.DbConnection()
}

func main() {
	r := gin.Default()
	Routers.User(r)
	r.Run()
}
