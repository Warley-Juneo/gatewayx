package api

import (
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "time"
    "net/url"
    "github.com/Warley-Juneo/gatewayx/core"
)

func HealthCheck(c *gin.Context) {
    c.JSON(200, gin.H{"status": "ok"})
}

func CreateProxyHandler(target *url.URL) gin.HandlerFunc {
    proxy := core.NewReverseProxy(target)
    return func(c *gin.Context) {
        proxy.ServeHTTP(c)
    }
}

func Login(c *gin.Context) {
    // Cria um token JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": "user123", // Subject (usuário)
        "exp": time.Now().Add(time.Hour * 1).Unix(), // Expira em 1 hora
    })

    // Assina o token com uma chave secreta
    tokenString, err := token.SignedString([]byte("secret_key"))
    if err != nil {
        c.JSON(500, gin.H{"error": "Erro ao gerar token"})
        return
    }

    c.JSON(200, gin.H{"token": tokenString})
}