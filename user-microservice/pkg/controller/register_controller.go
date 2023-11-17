package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"user-microservice/pkg/models"
	"user-microservice/pkg/repositories"
)

func Register(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *models.User
		var err error

		user = &models.User{}
		user.SellerID = c.GetInt("sellerId")
		user.Email = c.GetString("email")
		user.Password = c.GetString("password")

		if err = c.BindJSON(user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "malformed request"})
			return
		}

		userRepository := repositories.NewRepository(databaseConnection)
		user, err = userRepository.Create(user)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
