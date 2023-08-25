package main

import (
	"login-register-email/config"
	"login-register-email/routes"
)

func main() {
	config.InitializeDB()

	//init routes
	routes.InitRoutes()
}
