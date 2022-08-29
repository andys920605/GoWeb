package middlewares

import (
	srv "GoWeb/models/service"
	"GoWeb/service"
	service_interface "GoWeb/service/interface"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// ParseToken Parse token
func ParseToken(tokenString string) (*srv.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &srv.Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return service.JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// Valid token
	if claims, ok := token.Claims.(*srv.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware
func JWTAuthMiddleware(svc service_interface.ILoginSrv) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get token from Header.Authorization field.
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization is null in Header",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Format of Authorization is wrong",
			})
			c.Abort()
			return
		}
		// parts[0] is Bearer, parts is token.
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Token.",
			})
			c.Abort()
			return
		}
		// check cache
		if err := svc.CheckTokenExist(mc.Account); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err,
			})
			c.Abort()
			return
		}
		// Store Account info into Context
		c.Set("account", mc)
		// After that, we can get Account info from c.Get("account")
		c.Next()
	}
}
