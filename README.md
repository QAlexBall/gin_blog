# gin_blog
go gin framework example

### Usage
Swagger
```
go get -u github.com/swaggo/swag/cmd/swag
swag init

example
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
```

general Start
```shell
sudo /etc/init.d/postgresql restart
go run main.go
```

Use air (hot restart)
```
# create .air.conf in project folder
air -c .air.conf
```