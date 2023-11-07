package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type order struct {
	Amount     float64 `json:"amount" binding:"required,gte=0.0"`
	CustomerId int     `json:"customer_id"`
	Date       string  `json:"date" binding:"required"`
}

// curl -X POST -H 'Content-Type: application/json' -d '{"amount": 55.50, "date": "2023-11-05"}' http://localhost:8089/customers/32/orders
func CreateOrderForCustomer(c *gin.Context) {
	customerId, err := strconv.Atoi(c.Param("customer-id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var order order
	if err := c.ShouldBind(&order); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	order.CustomerId = customerId
	c.IndentedJSON(http.StatusOK, order)
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.POST("/customers/:customer-id/orders", CreateOrderForCustomer)

	if err := r.Run(":8089"); err != nil {
		panic(err)
	}
}
