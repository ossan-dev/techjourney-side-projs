package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl http://localhost:8089/safe
func SafeHandler(c *gin.Context) {
	keyToGet := "foo"
	c.Set("foo", "baz")
	val, exists := c.Get(keyToGet)
	if !exists {
		c.String(http.StatusInternalServerError, fmt.Sprintf("the key %q doesn't exist\n", keyToGet))
		return
	}
	c.String(http.StatusOK, val.(string))
}

// curl http://localhost:8089/unsafe
func UnsafeHandler(c *gin.Context) {
	keyToGet := "foo2"
	c.Set("foo", "baz")
	val := c.MustGet(keyToGet)
	c.String(http.StatusOK, val.(string))
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/safe", SafeHandler)
	r.GET("/unsafe", UnsafeHandler)

	if err := r.Run(":8089"); err != nil {
		panic(err)
	}
}
