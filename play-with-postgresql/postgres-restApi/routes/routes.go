package routes

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/praveenmahasena647/restApi/dbs"
)

func RunServer() *gin.Engine {
	var router *gin.Engine = gin.Default()

	router.GET("/", getAll)
	router.GET("/:id", getOne)
	router.POST("/", postOne)
	router.DELETE("/:id", deleteOne)
	router.DELETE("/", deleteAll)

	return router

}

func getAll(c *gin.Context) {
	var someThing = dbs.DB.Find(&dbs.Person{})
	log.Println(someThing)
}

func getOne(c *gin.Context) {

}

func postOne(c *gin.Context) {
	var data *dbs.Person = &dbs.Person{}

	c.BindJSON(data)
	c.JSONP(200, "cum")
}

func deleteOne(c *gin.Context) {

}

func deleteAll(c *gin.Context) {

}
