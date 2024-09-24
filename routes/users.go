package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	error := context.ShouldBind(&user)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the user"})
		return
	}
	error = user.Save()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the user"})
		return
	}
	context.JSON(http.StatusOK, user)
}

func login(context *gin.Context) {
	var user models.User
	error := context.ShouldBind(&user)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the user"})
		return
	}
	error = user.ValidateCredentials()
	if error != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": error.Error()})
		return
	}
	token, error := utils.GenerateToken(user.Email, user.ID)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
