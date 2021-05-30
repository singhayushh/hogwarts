package services

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB(DbName, DbUser, DbPass, DbType, DbHost, DbPort string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPass)
	DB, err = gorm.Open(DbType, DBURL)
	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	} else {
		fmt.Println("Connection to Postgres Database established.")
	}
	DB.Debug().AutoMigrate()
}
