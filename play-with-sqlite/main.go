package main

import (
	"log"
	"os"

	"github.com/praveenmahasena647/sqlite3/routers"
)

func main() {
	var err error = routers.RunServer()
	if err != nil {
		log.Println("server Error")
		os.Exit(0)
	}
}
