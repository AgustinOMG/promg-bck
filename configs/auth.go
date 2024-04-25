package configs

import (
	"context"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authClient *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetHeader("uid")
		token, err := authClient.VerifyIDToken(context.Background(), uid)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Set("uid", token.UID)
		c.Next()
	}
}
