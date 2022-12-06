package dbs

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
	uri string = "host=localhost user=postgres password=abc123 dbname=cum port=5432"
)

func init() {
	DB, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		log.Println("error during connecting to db")
		log.Println(err.Error())
		os.Exit(2)
	}
	var mErr error = DB.AutoMigrate(&Person{})
	if mErr != nil {
		log.Println("error during AutoMigrate")
		log.Println(mErr.Error())
	}
}
