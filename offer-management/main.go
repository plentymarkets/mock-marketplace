package main

import (
	"log"
	"offer-management/pkg/database"
	"offer-management/pkg/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	dsn := database.GetMariaDBDSN()
	con := database.CreateConnection(dsn)

	router.Offer(con, engine)
	router.User(con, engine)

	err := engine.Run(":3002")

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
