package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	valid := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {

	}
}
