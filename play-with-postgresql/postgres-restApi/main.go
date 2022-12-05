package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/praveenmahasena647/restApi/routes"
)

func main() {
	log.Println("cum")

	var router *gin.Engine = routes.RunServer()

	var err error = router.Run("localhost:42069")
	if err != nil {
		log.Println("Error during init server")
		log.Println(err.Error())
		os.Exit(22)
	}
}
