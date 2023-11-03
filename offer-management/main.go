package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"offer-management/pkg/database"
	"offer-management/pkg/router"
)

func main() {
	engine := gin.Default()

	dsn := database.GetMariaDBDSN()
	con := database.CreateConnection(dsn)

	router.Offer(con, engine)
	router.User(con, engine)
	router.Order(con, engine)

	err := engine.Run(":3002")

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
