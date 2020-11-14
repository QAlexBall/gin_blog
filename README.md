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

Database
```
sudo /etc/init.d/postgresql restart
```

General Running
```shell
go run main.go
```

Use air (hot restart)
```
# create .air.conf in project folder
air -c .air.conf
```

API
```
# get user list
curl -X GET "http://localhost:8080/api-user/user" -H "accept: application/json"

# create user
curl -H "Content-Type: application/json" \
--request POST \
--data '{"name":"chris","email":"chris@me.com","phone":"100","address":"nowhere"}'\

# get user by id
curl -X GET "http://localhost:8080/api-user/user/1" -H "accept: application/json"

# delete user
curl -X DELETE "http://localhost:8080/api-user/user/1" -H "accept: application/json"
```

## frontend
Yew

```
cargo web start
```