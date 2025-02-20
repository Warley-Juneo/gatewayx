package core

import (
    "github.com/golang-jwt/jwt/v5"
    "github.com/gin-gonic/gin"
    "net/http"
)

func JWTMiddleware(secret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        })
        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
            return
        }

        c.Next()
    }
}