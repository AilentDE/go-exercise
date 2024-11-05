package router

import (
	"net/http"
	"restful-api-sample/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
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
			"message": "Could not find this event information.",
		})
		return
	}

	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not register this event.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Register successfully.",
	})
}

func cancelRegisterForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
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
			"message": "Could not find this event information.",
		})
		return
	}

	err = event.CancelRegister(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not cancel register this event.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Cancel successfully.",
	})
}
