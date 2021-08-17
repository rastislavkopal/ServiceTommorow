package middleware

import (
	// "github.com/gin-contrib/sessions"
	"backend/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	userkey = "user"
)

// AuthRequired is a simple middleware to check the session
func AuthRequired(ctx *gin.Context) {
	err := service.TokenValid(ctx.Request)

	if err != nil {
		// ctx.AbortWithError(http.StatusForbidden, err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to perform this action"})
		return
	}
	// session = sessions.Default(ctx)

	// user := session.Get(userkey)

	// if user == nil {
	// 	// Abort the request with the appropriate error code
	// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	// 	return
	// }

	// Continue down the chain to handler etc
	ctx.Next()
}
