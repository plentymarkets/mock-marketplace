package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"product-management/pkg/database"
	"product-management/pkg/router"
)

func main() {

	engine := gin.Default()
	mariadb := database.NewMariaDBDatabase()

	router.Product(mariadb, engine)
	router.Variant(mariadb, engine)

	err := engine.Run(":3004")

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
