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

// check if selected workspace id is allowed for user with ID from JWT token
func WorkspaceUser(ctx *gin.Context) {
	err := service.TokenValid(ctx.Request)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to perform this action"})
		return
	}

	acccessDetails, error := service.ExtractTokenMetadata(ctx.Request)

	if error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Could not extract Token metadata"})
		return
	}

	if acccessDetails.UserId == 0 {
		// check if UserId is allowed for workspace with ":id"
	}

	// Continue down the chain to handler etc
	ctx.Next()
}
