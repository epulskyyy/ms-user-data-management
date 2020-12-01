package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"ms-user-data-management/models"
	"os"
)

var db *gorm.DB

func ConnectToDB() *gorm.DB {
	conn := os.Getenv("MYSQL_URL")
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		fmt.Println("[CONFIG.ConnectDB] error when connect to database")
		log.Fatal(err)
	}else {
		fmt.Println("SUCCES CONNECT TO DATABASE")
	}

	models.Migrate(db)

	return db
}
