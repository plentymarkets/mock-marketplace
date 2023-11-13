package main

import (
	"fmt"
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

	address := fmt.Sprintf("%s:%s",
		os.Getenv("HOST"),
		os.Getenv("PRODUCT_PORT"),
	)
	err := engine.Run(address)

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
