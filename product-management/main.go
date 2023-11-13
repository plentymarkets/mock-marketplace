package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"product-management/pkg/database"
	"product-management/pkg/router"
)

func main() {
	engine := gin.Default()

	dsn := database.GetMariaDBDSN()
	con := database.CreateConnection(dsn)

	router.Auth(con, engine)
	router.Product(con, engine)
	router.Variant(con, engine)

	err := engine.Run(os.Getenv("GIN_PORT"))

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
