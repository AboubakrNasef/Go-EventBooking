package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/ping", getEvents)

	engine.Run(":8080")
}

func getEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"events": []string{"event1", "event2", "event3"},
	})
}
