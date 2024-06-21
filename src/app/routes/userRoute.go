package routes

import (
	"github.com/gbubemi22/gowith_ginsetup/src/app/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userController := controllers.NewUserController()

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/user/create", userController.CreateUserHandler)
		}
	}
}
