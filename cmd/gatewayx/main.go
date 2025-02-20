package main

import (
    "log"
    "net/url"
    "strconv"
	"time"

    "github.com/Warley-Juneo/gatewayx/api"
    "github.com/Warley-Juneo/gatewayx/config"
    "github.com/Warley-Juneo/gatewayx/core"
    "github.com/Warley-Juneo/gatewayx/middleware"
    "github.com/Warley-Juneo/gatewayx/storage"
    "github.com/gin-gonic/gin"
)

func main() {
    cfg, err := config.LoadConfig("configs/config.yaml")
    if err != nil {
        log.Fatalf("Erro ao carregar configurações: %v", err)
    }

    // Redis para rate limiting
    redisClient := storage.NewRedisClient("localhost:6379")
    if err := storage.PingRedis(redisClient); err != nil {
        log.Fatalf("Erro ao conectar ao Redis: %v", err)
    }

    if !cfg.Server.Debug {
        gin.SetMode(gin.ReleaseMode)
    }
    router := gin.Default()
    router.Use(middleware.Logging())

    // Health Check
   // Health Check
	router.GET("/health", api.HealthCheck)

	// Endpoint de login para gerar tokens JWT
	router.POST("/login", api.Login)

	// Middleware de autenticação JWT
	router.Use(core.JWTMiddleware("secret_key"))

    // Middleware de rate limiting
    router.Use(middleware.RateLimiter(redisClient, 100, time.Minute))

    // Rotas dinâmicas
    for _, route := range cfg.Routes {
        targetURL, err := url.Parse(route.Target)
        if err != nil {
            log.Fatalf("URL inválida para a rota %s: %v", route.Name, err)
        }
        router.Any(route.Path, api.CreateProxyHandler(targetURL))
    }

    log.Printf("GatewayX iniciando na porta :%d", cfg.Server.Port)
    router.Run(":" + strconv.Itoa(cfg.Server.Port))
}