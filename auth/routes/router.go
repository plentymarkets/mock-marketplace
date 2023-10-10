package routes

import (
	"auth/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func RegisterRoutes() {
	// start default
	r := gin.Default()
	// main route /api
	api := r.Group("/api")
	// main route /api/auth and call func auth
	api.POST("/auth", controller.Auth())
	err := r.Run()

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
