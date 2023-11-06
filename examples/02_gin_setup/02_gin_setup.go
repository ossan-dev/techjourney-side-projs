package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl http://localhost:8089/example			=> GET
// curl http://localhost:8089/health			=> GET
// curl -X POST http://localhost:8089/health	=> POST

func HandleHealthCheck(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintln("the system is health"))
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/example", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintln("this is a demo endpoint"))
	})
	r.GET("/health", HandleHealthCheck)

	if err := r.Run(":8089"); err != nil {
		panic(err)
	}
}
