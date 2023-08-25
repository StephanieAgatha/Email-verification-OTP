package routes

import (
	"github.com/gin-gonic/gin"
	"login-register-email/controller"
	"login-register-email/middleware"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controller.Register)
	r.POST("/v1/auth/user/login", controller.HandleLogin)
	r.POST("/v1/auth/user/login/email_otp/start", controller.LoginStart)

	//middleware
	afterAuth := r.Group("/app")
	afterAuth.Use(middleware.AuthMiddleware())
	{
		afterAuth.GET("/books", controller.Books)
	}

	r.Run(":3000")
	return r
}
