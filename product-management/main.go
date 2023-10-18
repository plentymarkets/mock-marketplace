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

	router.ProductRouter(mariadb, engine)
	router.VariantRouter(mariadb, engine)

	err := engine.Run(":3004")

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
