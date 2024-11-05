package router

import (
	"net/http"
	"restful-api-sample/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events.",
		})
		return
	}

	if events == nil {
		events = make([]models.Event, 0)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	var err error

	err = ctx.ShouldBindBodyWithJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong body.",
		})
		return
	}

	userId := ctx.GetInt64("userId")
	event.UserId = userId

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create this event.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event created!",
		"event":   event,
	})
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
		})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event.",
		})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
		})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Coule not fetch event.",
		})
		return
	}
	if event.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update event.",
		})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data body.",
		})
		return
	}

	updatedEvent.Id = eventId
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update this event.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update successfully.",
	})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id.",
		})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Coule not fetch event.",
		})
		return
	}
	if event.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to delete event.",
		})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete this event.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete successfully.",
	})
}
