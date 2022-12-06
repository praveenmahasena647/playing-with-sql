package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/praveenmahasena647/restApi/dbs"
)

func RunServer() *gin.Engine {
	var router = gin.Default()

	router.GET("/", getAll)
	router.GET("/:id", getOne)
	router.POST("/", postOne)
	router.PUT("/:id", putOne)
	router.DELETE("/:id", deleteOne)
	router.DELETE("/", deleteAll)

	return router
}

func getAll(c *gin.Context) {
	var people []*dbs.Person
	var err = dbs.DB.Find(&people)
	if err.Error != nil {
		log.Println("Error During getting all")
		c.JSONP(404, "Error During getting all")
		return
	}
	c.JSONP(200, people)
}

func getOne(c *gin.Context) {
	var id string = c.Param("id")
	var person = &dbs.Person{}
	var err = dbs.DB.Find(&person, id)
	if err.Error != nil {
		log.Println("Error During getting one")
		c.JSONP(404, "Error During getting one")
		return
	}
	c.JSONP(200, &person)
}
func postOne(c *gin.Context) {
	var person *dbs.Person = &dbs.Person{}
	c.BindJSON(&person)
	var err = dbs.DB.Create(&person)
	if err.Error != nil {
		log.Println("Error During Inserting one")
		c.JSONP(404, "Error During Inserting one")
		return
	}
	c.Status(200)
}
func putOne(c *gin.Context) {
	var id string = c.Param("id")
	var person *dbs.Person = &dbs.Person{}
	var err = dbs.DB.First(&person, id)
	if err.Error != nil {
		log.Println("Error During Put Event one")
		c.JSONP(404, "Error During Put Event one")
		return
	}
	c.BindJSON(&person)
	var erro = dbs.DB.Save(&person)
	if erro.Error != nil {
		log.Println("Error During Put Event one")
		c.JSONP(404, "Error During Put Event one")
		return
	}
}
func deleteOne(c *gin.Context) {
	var id string = c.Param("id")
	var err = dbs.DB.Delete(&dbs.Person{}, id)
	if err.Error != nil {
		log.Println("Error During delete Event one")
		c.JSONP(404, "Error During delete Event one")
		return
	}
	c.Status(200)
}

func deleteAll(c *gin.Context) {
	var err = dbs.DB.Delete(&dbs.Person{})
	if err.Error != nil {
		log.Println("Error During delete Event ")
		c.JSONP(404, "Error During delete Event ")
		return
	}
	c.Status(200)
}
