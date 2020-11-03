package routes

import (
	"gin_blog/controllers"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gin_blog/docs"

	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.New()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	userAPI := r.Group("api-user")
	{
		userAPI.GET("user", controllers.GetUsers)
		userAPI.POST("user", controllers.CreateUser)
		userAPI.GET("user/:id", controllers.GetUserByID)
		userAPI.PUT("user/:id", controllers.UpdateUser)
		userAPI.DELETE("user/:id", controllers.DeleteUser)
	}
	return r
}
