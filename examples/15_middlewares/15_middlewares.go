package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl http://localhost:8089/example
func HandleExample(c *gin.Context) {
	fmt.Println("HandleExample()")
	c.String(http.StatusOK, fmt.Sprintln("HandleExample..."))
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		fmt.Println("BEFORE - HandleExample")
		c.Next()
		fmt.Println("AFTER - HandleExample")
	})

	r.GET("/example", HandleExample)

	if err := r.Run(":8089"); err != nil {
		panic(err)
	}
}
