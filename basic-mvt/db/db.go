package db

import (
	"log"
	"os"

	"github.com/michaeljs1990/sqlitestore"
	"github.com/tjgurwara99/echo-boilerplates/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

var SessionStore *sqlitestore.SqliteStore

func init() {
	var err error
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}
	dbLoc := env + ".db"
	DB, err = gorm.Open(sqlite.Open(dbLoc), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Add models here for migrations - gorm will create tables for them
	err = DB.AutoMigrate(&models.User{}, &models.Permission{}, &models.Company{}, &models.Role{})
	if err != nil {
		log.Fatal(err)
	}
	SessionStore, err = sqlitestore.NewSqliteStore(dbLoc, "sessions", "/", 3600, []byte("secret"))
	if err != nil {
		log.Fatal(err)
	}
	SessionStore.Options.HttpOnly = true
	SessionStore.Options.Secure = true
}
