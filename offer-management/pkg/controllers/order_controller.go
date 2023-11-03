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

const OrdersPerPage = 10

type OrderController struct {
	orderRepository repositories.OrderRepositoryContract
}

func NewOrderController(orderRepository *repositories.OrderRepository) OrderController {
	return OrderController{
		orderRepository: orderRepository,
	}
}

func (controller *OrderController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		pageStr := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(pageStr)

		if err != nil {
			log.Printf("Invalid page number unsupported format %s", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid page number format! Page number should be an integer."})
			return
		}

		orders, pageCount, err := controller.orderRepository.FetchAll(page, OrdersPerPage)

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
			"data":      orders,
			"pageCount": pageCount,
		})
		c.Done()
	}
}

func (controller *OrderController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var order = models.Order{}
		err := c.BindJSON(&order)

		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		order, err = controller.orderRepository.Create(order)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"message": "Success",
			"data":    order,
		})
		c.Done()
	}
}
