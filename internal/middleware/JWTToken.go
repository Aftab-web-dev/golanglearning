package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtkey = []byte(os.Getenv("JWT_SECRET"))


func GenerateToken(userID string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    })
    return token.SignedString(jwtkey)
}


func AuthMiddleware( c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		c.Abort()
		return
	}

	 tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	 token , err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtkey , nil
	 })
	 
	 if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Invaild token"})
		c.Abort()
		return
	 }
	 c.Next()
}