package utils

import (
	"encoding/base64"
	"log"
	"time"

	"github.com/otisnado/nepackage/models"
)

func AdminCreation() {

	var adminUser models.User
	adminUser.Name = "Admin"
	adminUser.Lastname = "System"
	adminUser.Mail = "admin@nepackage.org"
	adminUser.Username = "root"
	adminUser.Password = "admin"
	adminUser.Role = "1"
	adminUser.Group = 1
	adminUser.State = true

	if err := CheckIfMailExists(adminUser.Mail); err != nil {
		return
	}

	if err := CheckIfUsernameExists(adminUser.Username); err != nil {
		return
	}

	passwordOutput := base64.StdEncoding.EncodeToString([]byte(adminUser.Password))
	log.Println("Password encoded in base64: ", passwordOutput)

	if err := adminUser.HashPassword(adminUser.Password); err != nil {
		return
	}

	user := models.User{Name: adminUser.Name, Lastname: adminUser.Lastname, Username: adminUser.Username, Mail: adminUser.Mail, Password: adminUser.Password, Role: adminUser.Role, Group: adminUser.Group, State: adminUser.State, CreatedAt: time.Now(), UpdatedAt: adminUser.UpdatedAt}
	models.DB.Create(&user)
}

func GroupAdminCreation() {
	var groupAdmin models.Group
	groupAdmin.Name = "Administrators"
	result := models.DB.Where("name = ?", groupAdmin.Name).Find(&groupAdmin)
	if result.RowsAffected > 0 {
		return
	}
	group := models.Group{Name: groupAdmin.Name, CreatedAt: time.Now(), UpdatedAt: groupAdmin.UpdatedAt}
	models.DB.Create(&group)

}

func RoleAdminCreation() {
	var roleAdmin models.Role
	roleAdmin.Name = "Administrators"
	result := models.DB.Where("name = ?", roleAdmin.Name).Find(&roleAdmin)
	if result.RowsAffected > 0 {
		return
	}
	role := models.Role{Name: roleAdmin.Name}
	models.DB.Create(&role)
}

func PolicyAdminCreation() {
	var policyAdmin models.Policy
	policyAdmin.Name = "Administrators"
	policyAdmin.Path = "/*"
	policyAdmin.AuthorizedRole = 1
	result := models.DB.Where("name = ?", policyAdmin.Name).Find(&policyAdmin)
	if result.RowsAffected > 0 {
		return
	}
	policy := models.Policy{Name: policyAdmin.Name, Path: policyAdmin.Path, AuthorizedRole: policyAdmin.AuthorizedRole}
	models.DB.Create(&policy)
}
