package database

import (
	"log"
	"os"

	"github.com/nepackage/nepackage/models"
	"github.com/nepackage/nepackage/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	if os.Getenv("DSN") != "" || utils.EnvExist("DSN") {
		dsn := os.Getenv("DSN")
		log.Println(dsn)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}

		err = db.AutoMigrate(models.Project{}, models.Language{}, models.User{}, models.Group{}, models.Role{}, models.Policy{}, models.GithubCredential{}, models.Provider{})
		if err != nil {
			return
		}
		DB = db
	} else {
		db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}

		err = db.AutoMigrate(models.Project{}, models.Language{}, models.User{}, models.Group{}, models.Role{}, models.Policy{}, models.GithubCredential{}, models.Provider{})
		if err != nil {
			return
		}
		DB = db
	}
}
