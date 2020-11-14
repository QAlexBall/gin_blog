package controllers

import (
	"fmt"
	"gin_blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers => get all user
// @Summary Get User List
// @Produce  json
// @Success 200 {} string
// @Failure 500 {} string
// @Router /api-user/user [get]
func GetUsers(c *gin.Context) {
	var user []models.User
	err := models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//CreateUser ... Create User
// @Summary Create User
// @Produce json
// @Success 200 {} string
// @Failure 500 {} string
// @Router /api-user/user [post]
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByID ... Get the user by id
// @Summary Get User By ID
// @Produce json
// @Success 200 {} string
// @Failure 500 {} string
// @Router /api-user/user/:id [get]
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
// @Summary Update User
// @Produce json
// @Success 200 {} string
// @Failure 500 {} string
// @Router /api-user/user/:id [put]
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
// @Summary Delete User
// @Produce json
// @Success 200 {} string
// @Failure 500 {} string
// @Router /api-user/user/:id [delete]
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
