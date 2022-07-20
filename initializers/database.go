package initializers

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "root@tcp(127.0.0.1:3306)/ordermgmt"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// var err error
	// dsn := os.Getenv("DB_STRING")
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
