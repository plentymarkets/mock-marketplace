package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"offer-management/pkg/controllers"
	"offer-management/pkg/middlewares"
	"offer-management/pkg/repositories"
)

func User(mariadb *gorm.DB, engine *gin.Engine) {

	// Homework = Check what the Engin
	userRepository, _ := repositories.NewUserRepository(mariadb)
	userController := controllers.NewUserController(userRepository)

	user := engine.Group("/api/user").Use(middlewares.Authenticate())

	user.GET("/", userController.GetAll())
	user.GET("/:id", userController.GetByID())
	user.POST("/login", userController.Login())
	// user.GET("/:user_name", userController.GetByName())
	user.POST("/", userController.Create())
	user.PUT("/", userController.Update())
	user.DELETE("/:id", userController.Delete())
	// user.DELETE("/:user_name", userController.Delete())
}
