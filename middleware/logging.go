package middleware

import (
    "log"
    "time"
    "github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        logData := map[string]interface{}{
            "method":   c.Request.Method,
            "path":     c.Request.URL.Path,
            "status":   c.Writer.Status(),
            "duration": time.Since(start).String(),
            "clientIP": c.ClientIP(),
        }
        log.Printf("%+v\n", logData)
    }
}