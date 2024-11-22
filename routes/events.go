package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

// Logic for fetching an Events list
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again laster."})
		return
	}
	context.JSON(http.StatusOK, events)
}

// Logic for fetching a single Event by ID
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again laster."})
		return
	}

	context.JSON(http.StatusOK, event)
}

// Logic for creating an Event
func createPost(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong!!", "error": err})
		return
	}

	event.UserId = context.GetInt64("userId")

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again laster."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

// Logic for upading an Event
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id."})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong!!", "error": err})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event.", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted successfully"})
}
