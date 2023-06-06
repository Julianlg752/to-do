package middleware

import (
	"net/http"
	"strings"
	"todo/config"

	"github.com/ansel1/merry/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, merry.New("Invalid"))
			return
		}

		token, err := jwt.Parse(t[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.AbortWithStatusJSON(http.StatusBadRequest, merry.New("Invalid Signature"))
			}
			return config.Secret(), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, merry.New("Internal Server Error"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, merry.New("Invalid Token"))
			return
		}
		c.Set("x-user-id", claims["userId"].(float64))
		c.Next()
	}
}
