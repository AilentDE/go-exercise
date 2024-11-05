package router

import (
	"fmt"
	"net/http"
	"restful-api-sample/models"
	"restful-api-sample/utils"

	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	var err error

	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong body.",
		})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create this user.",
		})
		return
	}

	fmt.Println(user)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created!",
		"user": map[string]any{
			"id":    user.Id,
			"email": user.Email,
		},
	})
}

func login(ctx *gin.Context) {
	var user models.User
	var err error

	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong body.",
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not authenticate user.",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error with quth info.",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successfully!",
		"token":   token,
	})
}
