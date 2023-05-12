package main

import (
	"github.com/nepackage/nepackage/config"
	"github.com/nepackage/nepackage/database"
	"github.com/nepackage/nepackage/routes"
)

func main() {
	database.ConnectDatabase()
	config.RoleAdminCreation()
	config.PolicyAdminCreation()
	config.GroupAdminCreation()
	config.AdminCreation()
	router := routes.InitRouter()
	router.Run()

}
