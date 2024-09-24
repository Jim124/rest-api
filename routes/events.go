package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, error := models.GetAllEvents()
	if error != nil {
		fmt.Println(error)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch data from db."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	fmt.Println(userId)
	var event models.Event
	error := context.ShouldBind(&event)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request"})
		return
	}

	event.UserID = userId
	event.DateTime = time.Now()
	error = event.Save()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "save event error"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func getSingleEvent(context *gin.Context) {
	var paramId = context.Param("id")
	id, error := strconv.ParseInt(paramId, 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	event, error := models.GetSingleEvent(id)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch an event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the id"})
		return
	}
	event, err := models.GetSingleEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}
	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "No authorized to update"})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBind(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request"})
		return
	}
	updatedEvent.ID = eventId
	updatedEvent.DateTime = time.Now()
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not update the request"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the id"})
		return
	}
	event, err := models.GetSingleEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}
	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Not authorized to delete the event"})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event delete successfully"})
}
