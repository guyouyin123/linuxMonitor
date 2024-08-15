package main

import (
	"github.com/gin-gonic/gin"
	"linuxMonitor/src"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/linuxInfo", src.Run)
}
