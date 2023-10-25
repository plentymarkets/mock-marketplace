package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"product-management/pkg/repositories"
)

type Person struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type AuthenticateController struct {
	userRepository repositories.UserRepositoryContract
}

func NewAuthenticateController(userRepository repositories.UserRepositoryContract) AuthenticateController {
	return AuthenticateController{
		userRepository: userRepository,
	}
}

func (controller *AuthenticateController) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var person = Person{}
		err := c.BindJSON(&person)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
			return
		}

		uuid := mdHashing(person.Username)

		// Create request to auth
		token := "Update_token_with_one_retrieved_from_Auth"

		user, err := controller.userRepository.FetchByID(uuid)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		user.UUID = uuid
		user.Token = token

		if user.ID == 0 {
			user, err = controller.userRepository.Create(user)
		} else {
			user, err = controller.userRepository.Update(user)
		}

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Message": "The user has been updated successfully",
		})
		c.Done()
	}
}

func mdHashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:]) // create a slice from an array
}
