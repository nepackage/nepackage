package config

import (
	"encoding/base64"
	"log"
	"time"

	"github.com/nepackage/nepackage/database"
	"github.com/nepackage/nepackage/models"
)

func AdminCreation() {

	var adminUser models.UserCreate
	adminUser.Name = "Admin"
	adminUser.Lastname = "System"
	adminUser.Mail = "admin@nepackage.org"
	adminUser.Username = "root"
	adminUser.Password = "admin"
	adminUser.Role = "1"
	adminUser.Group = 1
	adminUser.State = true

	passwordOutput := base64.StdEncoding.EncodeToString([]byte(adminUser.Password))
	log.Println("Password encoded in base64: ", passwordOutput)

	if err := adminUser.HashPassword(adminUser.Password); err != nil {
		return
	}

	user := models.User{Name: adminUser.Name, Lastname: adminUser.Lastname, Username: adminUser.Username, Mail: adminUser.Mail, Password: adminUser.Password, Role: adminUser.Role, Group: adminUser.Group, State: adminUser.State, CreatedAt: time.Now(), UpdatedAt: adminUser.UpdatedAt}
	database.DB.Create(&user)
}

func GroupAdminCreation() {
	var groupAdmin models.Group
	groupAdmin.Name = "Administrators"
	result := database.DB.Where("name = ?", groupAdmin.Name).Find(&groupAdmin)
	if result.RowsAffected > 0 {
		return
	}
	group := models.Group{Name: groupAdmin.Name, CreatedAt: time.Now(), UpdatedAt: groupAdmin.UpdatedAt}
	database.DB.Create(&group)

}

func RoleAdminCreation() {
	var roleAdmin models.Role
	roleAdmin.Name = "Administrators"
	roleAdmin.Policies = "1"
	result := database.DB.Where("name = ?", roleAdmin.Name).Find(&roleAdmin)
	if result.RowsAffected > 0 {
		return
	}
	role := models.Role{Name: roleAdmin.Name, Policies: roleAdmin.Policies}
	database.DB.Create(&role)
}

func PolicyAdminCreation() {
	var policyAdmin models.Policy
	policyAdmin.Name = "Administrators"
	policyAdmin.Path = "/*"
	policyAdmin.AuthorizedMethods = "GET,POST,PATCH,UPDATE,DELETE"
	result := database.DB.Where("name = ?", policyAdmin.Name).Find(&policyAdmin)
	if result.RowsAffected > 0 {
		return
	}
	policy := models.Policy{Name: policyAdmin.Name, Path: policyAdmin.Path, AuthorizedMethods: policyAdmin.AuthorizedMethods}
	database.DB.Create(&policy)
}
