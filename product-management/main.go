package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"product-management/pkg/database"
	"product-management/pkg/router"
	"product-management/pkg/utils/env-handler"
)

func main() {
	env_handler.LoadEnvironment()

	engine := gin.Default()

	dsn := database.GetMariaDBDSN()
	con := database.CreateConnection(dsn)

	router.Product(con, engine)
	router.Variant(con, engine)

	address := fmt.Sprintf("%s:%s",
		os.Getenv("GIN_HOST"),
		os.Getenv("GIN_PORT"),
	)
	err := engine.Run(address)

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
