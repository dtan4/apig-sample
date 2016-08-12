package main

import (
	"github.com/dtan4/apig-sample/db"
	"github.com/dtan4/apig-sample/server"
)

// main ...
func main() {
	database := db.Connect()
	s := server.Setup(database)
	s.Run(":8080")
}
