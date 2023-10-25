package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"product-management/pkg/models"
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

		// TODO - Create request to auth
		// token, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
		token := "Update_token_with_one_retrieved_from_Auth"

		uuid := mdHashing(person.Username)
		user := models.User{UUID: uuid}
		user, err = controller.userRepository.FetchByUser(user)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		user.UUID = uuid
		user.Token = token

		message := ""

		if user.ID == 0 {
			user, err = controller.userRepository.Create(user)
			message = "The user has been registered successfully"
		} else {
			user, err = controller.userRepository.Update(user)
			message = "The user has been updated successfully"
		}

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Message": message,
		})
		c.Done()
	}
}

func mdHashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:]) // create a slice from an array
}
