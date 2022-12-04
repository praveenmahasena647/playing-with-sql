package main

import (
	"log"
	"os"

	"github.com/praveenmahasena647/postgresql/routes"
)

func main() {
	log.Println("fuck")

	var router = routes.RunServer()

	var err error = router.Run("localhost:42069")

	if err != nil {
		log.Println("Server Cannot Start")
		os.Exit(2)
	}
}
