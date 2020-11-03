package main

import (
	"gin_blog/app"
	"gin_blog/models"
	"gin_blog/routes"
)

func initialize() {
	app.Setup()
}

func main() {
	initialize()

	app.DB.AutoMigrate(&models.User{})
	r := routes.SetupRouter()
	r.Run()
}
