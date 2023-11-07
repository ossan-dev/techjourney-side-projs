package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type creditCard struct {
	Number       string     `json:"number"`
	SecurityCode string     `json:"-"`
	Expiration   *time.Time `json:"expiration,omitempty"`
}

// curl http://localhost:8089/credit-cards
func GetCreditCards(c *gin.Context) {
	expirationDate := time.Now().AddDate(1, 0, 0)
	cards := []creditCard{
		{Number: "1234", SecurityCode: "647", Expiration: &expirationDate},
		{Number: "0000", SecurityCode: "012"},
	}
	c.IndentedJSON(http.StatusOK, cards)
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/credit-cards", GetCreditCards)

	if err := r.Run(":8089"); err != nil {
		panic(err)
	}
}
