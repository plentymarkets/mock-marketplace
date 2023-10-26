package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"offer-management/pkg/models"
	"offer-management/pkg/repositories"
	"strconv"
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

func (controller *UserController) GetByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("user_name")
		_, err := strconv.Atoi(username)

		if err != nil {
			log.Printf("User not found: %s", err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}

		user, err := controller.userRepository.FetchByName(username)

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

func (controller *UserController) Update() gin.HandlerFunc {
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

func (controller *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.PostForm("id")
		_, err := strconv.Atoi(id)

		if err = c.ShouldBind(&id); err != nil {
			log.Printf("Invalid request: %s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request: Check your input data"})
			return
		}

		fetchedUser, err := controller.userRepository.FetchByID(id)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		if fetchedUser.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
			return
		}

		providedName := c.PostForm("user_name")
		storedName := fetchedUser.UserName

		if providedName != storedName {
			log.Println("Invalid credentials")
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
			return
		}

		providedPassword := c.PostForm("user_password")
		storedPassword := fetchedUser.UserPassword

		if providedPassword != storedPassword {
			log.Println("Invalid credentials")
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
			return
		}

		token := "JWT-SECRET"
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
