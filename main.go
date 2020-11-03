package main

import (
	"fmt"
	"gin_blog/config"
	"gin_blog/models"
	"gin_blog/routes"
)

func main() {
	// endless.DefaultReadTimeOut =
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Status:", err)
	}
	db.AutoMigrate(&models.User{})

	r := routes.SetupRouter()
	r.Run()
}
