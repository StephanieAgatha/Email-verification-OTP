package main

import (
	"github.com/gin-gonic/gin"
	"login-register-email/config"
	"login-register-email/controller"
)

func main() {
	config.InitializeDB()

	r := gin.Default()
	r.POST("/register", controller.Register)
	r.POST("/v1/auth/user/login", controller.HandleLogin)
	r.POST("/v1/auth/user/login/email_otp/start", controller.LoginStart)

	r.Run(":3000")
}
