package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ndcuongg/event-booking-api.git/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get all events!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Get events successful!", "events": events})
}

func createEvent(context *gin.Context) {
	// Authentication Required

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request data!"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't save event. Try it later!"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't fetch event id."})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't found event id."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Fetched event successful", "event": event})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't fetch event id."})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch the event."})
		return
	}
	if userId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorization to update event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request data!"})
		return
	}
	updatedEvent.ID = id
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't update event!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Updated event successful"})

}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't fetch event id."})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch the event."})
		return
	}
	if userId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorization to update event"})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't delete the event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Deleted event successful"})
}
