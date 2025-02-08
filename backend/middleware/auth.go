package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("your_secret_key")

// Auth middleware verifies JWT token and sets the user claims in context.
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}

// RoleAuth middleware checks if the user's role is among allowed roles.
func RoleAuth(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}
		claims, ok := userClaims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Role not found in token"})
			return
		}

		authorized := false
		for _, allowed := range allowedRoles {
			if role == allowed {
				authorized = true
				break
			}
		}

		if !authorized {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied: insufficient role privileges"})
			return
		}

		c.Next()
	}
}
