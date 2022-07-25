package initializers

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "root@tcp(127.0.0.1:3306)/ordermgmt?parseTime=true"
	DB, err = gorm.Open("mysql", dsn)
	//defer DB.Close()

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
