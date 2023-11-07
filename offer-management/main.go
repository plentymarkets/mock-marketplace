package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"offer-management/pkg/database"
	"offer-management/pkg/router"
	"os"
)

func main() {
	engine := gin.Default()

	dsn := database.GetMariaDBDSN()
	con := database.CreateConnection(dsn)

	router.Auth(con, engine)
	router.Offer(con, engine)

	address := fmt.Sprintf("%s:%s",
		os.Getenv("HOST"),
		os.Getenv("OFFER_PORT"),
	)
	err := engine.Run(address)

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
