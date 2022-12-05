package dbs

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name  string `json:"name"`
	Nname string `json:"nname"`
	Age   uint   `json:"age"`
}
