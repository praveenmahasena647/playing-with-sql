package dbs

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Err error
)

type Person struct {
	gorm.Model
	Name  string `json:"name"`
	DName string `json:"dname"`
}

func init() {
	DB, Err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if Err != nil {
		log.Println("db Connection Err")
		log.Println(Err)
		os.Exit(0)
	}

	var migRateErr = DB.AutoMigrate(&Person{})
	if migRateErr != nil {
		log.Println("migRateErr")
		log.Println("Error")
	}

}
