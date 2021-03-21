package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
		return
	})

	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatalf("Invalid r run err:%s", err.Error())
	}
}

