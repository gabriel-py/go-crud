package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/amopromo/gochip/repository"
	"github.com/gin-gonic/gin"
)

func ValidateToken(repo repository.ApiClientRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")

		// Extract Bearer token
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			return
		}
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		apiKey := strings.TrimSpace(splitToken[1])

		// Validate the token
		apiClient, err := repo.Get(apiKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Get tenantID from apiClientID
		tenantID, err := repo.GetTenantIDByApiClientID(apiClient.ID)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Set tenantID in context
		c.Set("tenantID", tenantID)
		fmt.Println(c.Params)

		c.Next()
	}
}
