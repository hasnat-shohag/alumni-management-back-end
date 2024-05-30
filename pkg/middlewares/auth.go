package middlewares

import (
	"alumni-management-server/pkg/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ValidateToken validates the token in the Authorization header.
func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, "Authorization header missing")
		}

		// Extract the token from the header (Bearer <token>)
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			return c.JSON(http.StatusUnauthorized, "Invalid Authorization header format")
		}

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenParts[1], &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.LocalConfig.JwtSecret), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			// Set the user information in the context for further use
			c.Set("username", claims.Subject)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, "Invalid token")
	}

}
