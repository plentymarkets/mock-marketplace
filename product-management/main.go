package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"product-management/pkg/database"
	"product-management/pkg/router"
)

func main() {
	engine := gin.Default()

	dsn := database.GetMariaDBDSN()
	con := database.CreateConnection(dsn)

	router.Product(con, engine)
	router.Variant(con, engine)

	err := engine.Run(":3004")

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
