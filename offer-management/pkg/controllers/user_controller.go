package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"offer-management/pkg/models"
	"offer-management/pkg/repositories"
	"strconv"
	"time"
)

const UsersPerPage = 10

type UserController struct {
	userRepository repositories.UserRepositoryContract
}

func NewUserController(userRepository repositories.UserRepositoryContract) UserController {
	return UserController{
		userRepository: userRepository,
	}
}

func (controller *UserController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		pageStr := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(pageStr)

		if err != nil {
			log.Printf("Invalid page number unsupported format %s", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid page number format! Page number should be an integer."})
			return
		}

		users, pageCount, err := controller.userRepository.FetchAll(page, UsersPerPage)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		if page < 1 || page > pageCount {
			log.Println("Invalid page number!")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid page number! Please selet a page from 1 to %d", pageCount)})
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"data":      users,
			"pageCount": pageCount,
		})
		c.Done()
	}
}

func (controller *UserController) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		_, err := strconv.Atoi(id)

		if err != nil {
			log.Printf("Invalid user ID %s", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request: OfferID should be an integer value"})
			return
		}

		user, err := controller.userRepository.FetchByID(id)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		if user.ID == 0 {
			c.JSON(http.StatusOK, map[string]any{
				"data": nil,
			})
		} else {
			c.JSON(http.StatusOK, map[string]any{
				"data": user,
			})
		}
		c.Done()
	}
}

func (controller *UserController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user = models.User{}
		err := c.BindJSON(&user)

		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		user, err = controller.userRepository.Create(user)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"message": "Success",
			"data":    user,
		})
		c.Done()
	}
}

func (controller *UserController) Update() gin.HandlerFunc { // todo - investigate changes on the variant when changing the user.
	return func(c *gin.Context) {
		var user = models.User{}
		err := c.BindJSON(&user)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		user, err = controller.userRepository.Update(user)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		time.Sleep(100)

		c.JSON(http.StatusOK, map[string]any{
			"message": "User updated successfully",
			"data":    user,
		})
		c.Done()
	}
}

func (controller *UserController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		user, err := controller.userRepository.FetchByID(id)
		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		user.Deleted = true

		user, err = controller.userRepository.Update(user)
		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"message": "User with id " + id + " Has been deleted successfully",
			"data":    user,
		})
		c.Done()
	}
}
