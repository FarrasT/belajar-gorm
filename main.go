package main

import (
	"belajar-gorm/database"
	"belajar-gorm/router"
)

func main() {
	database.StartDB()
	var PORT = ":8080"

	router.StartServer().Run(PORT)

}
