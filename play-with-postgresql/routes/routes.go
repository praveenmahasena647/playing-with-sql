package routes

import "github.com/gin-gonic/gin"

func RunServer() *gin.Engine {
	var router = gin.Default()
	router.GET("/", getAll)
	router.GET("/:id", getOne)
	router.POST("/", postOne)
	router.DELETE("/:id", deleteOne)
	router.DELETE("/", deleteAll)
	return router
}

func getAll(c *gin.Context) {
	c.JSONP(200, "cum")
}

func getOne(c *gin.Context) {
	c.JSONP(200, "cum")
}

func postOne(c *gin.Context) {
	c.JSONP(200, "cum")
}

func deleteOne(c *gin.Context) {
	c.JSONP(200, "cum")
}

func deleteAll(c *gin.Context) {
	c.JSONP(200, "cum")
}
