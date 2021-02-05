package main

import (
	mysqldb "go_products_pagination/connection"
	"go_products_pagination/routes"
	"log"
)

func main() {

	db := mysqldb.SetupDB()

	r := routes.SetupRoutes(db)

	log.Println("listening on http://localhost:8000")

	r.Run(":8000")

}
