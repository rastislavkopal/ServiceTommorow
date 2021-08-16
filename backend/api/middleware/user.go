package middleware

import (
	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	userkey = "user"
)

// AuthRequired is a simple middleware to check the session
// TODO
func AuthRequired(ctx *gin.Context) {
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
