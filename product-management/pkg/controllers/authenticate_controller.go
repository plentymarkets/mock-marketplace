package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"product-management/pkg/client"
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

		response, err := client.Authenticate(person.Username, person.Password)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot Authenticate"})
			return
		}

		if response.StatusCode != http.StatusOK {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot Authenticate"})
			return
		}

		body, err := io.ReadAll(response.Body)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Internal server error "})
			return
		}

		uuid := mdHashing(person.Username)
		user := models.User{UUID: uuid}
		user, err = controller.userRepository.FetchByUser(user)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		user.UUID = uuid
		err = json.Unmarshal(body, &user)
		if err != nil {
			return
		}

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
			"token": user.Token,
		})
		c.Done()
	}
}

func mdHashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:]) // create a slice from an array
}
