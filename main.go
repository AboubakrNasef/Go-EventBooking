package main

import (
	"net/http"

	"eventBook.sample.com/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the Event Booking API")
	})

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")

}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents() // This should be replaced with actual event fetching logic
	c.JSON(http.StatusOK, events)
}
func createEvent(c *gin.Context) {
	var newEvent models.Event
	err := c.ShouldBindJSON(&newEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newEvent.ID = 1
	newEvent.UserID = 1
	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": newEvent})
}
