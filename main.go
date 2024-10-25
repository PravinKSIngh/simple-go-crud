package main

import (
	"simple-go-crud/db"
	"simple-go-crud/routes"
)

func main() {
	database := db.InitDB()
	defer database.Close()

	router := routes.SetupRouter(database)

	router.Run(":8080")
}
