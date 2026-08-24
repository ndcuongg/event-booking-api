package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ndcuongg/event-booking-api.git/models"
	"github.com/ndcuongg/event-booking-api.git/utils"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request data!"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't save user!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Saved user successful!"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request data!"})
		return
	}
	err = user.ValidateCredential()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Couldn't authenticate user!"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't generate token for user!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
