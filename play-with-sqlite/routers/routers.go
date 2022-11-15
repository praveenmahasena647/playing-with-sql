package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/praveenmahasena647/sqlite3/dbs"
)

func RunServer() error {
	var router = gin.Default()

	router.GET("/", getAll)
	router.GET("/:id", getOne)
	router.POST("/", insertOne)
	router.PUT("/:id", editOne)
	router.DELETE("/:id", deleteOne)
	return router.Run("localhost:42069")
}

func getAll(c *gin.Context) {
	var people []*dbs.Person
	var err = dbs.DB.Find(&people)
	if err.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSONP(200, people)
}

func getOne(c *gin.Context) {
	var id string = c.Param("id")
	var person *dbs.Person

	var err = dbs.DB.Find(&person, id)
	if err.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSONP(200, person)
}

func insertOne(c *gin.Context) {
	var person *dbs.Person
	c.BindJSON(&person)
	var done = dbs.DB.Create(&person)
	if done.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(200)
}

func editOne(c *gin.Context) {
	var id string = c.Param("id")
	var person *dbs.Person
	c.BindJSON(&person)
	var err = dbs.DB.Where("id = ?", id).Find(&dbs.Person{}, id).Updates(&person)

	if err.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	log.Println(err.Error)
}

func deleteOne(c *gin.Context) {
	var id string = c.Param("id")

	var err = dbs.DB.Delete(&dbs.Person{}, id)

	if err.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(200)
}
