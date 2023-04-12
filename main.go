package main

import (
	"github.com/nepackage/nepackage/models"
	"github.com/nepackage/nepackage/routes"
	"github.com/nepackage/nepackage/utils"
)

func main() {
	models.ConnectDatabase()
	utils.RoleAdminCreation()
	utils.PolicyAdminCreation()
	utils.GroupAdminCreation()
	utils.AdminCreation()
	router := routes.InitRouter()
	router.Run()

}
