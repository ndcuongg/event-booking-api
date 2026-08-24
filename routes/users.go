package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ndcuongg/event-booking-api.git/models"
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
