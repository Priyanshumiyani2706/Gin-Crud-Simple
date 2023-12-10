package main

import (
	"gin-crud/Config"
	Model "gin-crud/Models"
)

func init() {
	Config.DbConnection()
	Config.LoadEnv()
}

func main() {
	Config.DB.AutoMigrate(&Model.User{})
}
