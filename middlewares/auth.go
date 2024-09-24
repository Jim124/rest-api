package middlewares

import (
	"net/http"
	"strings"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	jwtToken := strings.Split(token, " ")
	if len(jwtToken) != 2 {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrectly formatted authorization header"})
		return
	}
	userId, error := utils.VerifyToken(jwtToken[1])
	if error != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	context.Set("userId", userId)
	context.Next()

}
