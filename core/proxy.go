package core

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

type ProxyService interface {
	ServeHTTP(c *gin.Context)
	SetTimeout(timeout time.Duration)
}

type ReverseProxy struct {
	target *url.URL
	client *http.Client
	cb     *gobreaker.CircuitBreaker
}

func NewReverseProxy(target *url.URL) ProxyService {
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "Proxy",
		Timeout: 5 * time.Second,
	})
	return &ReverseProxy{
		target: target,
		client: &http.Client{Timeout: 5 * time.Second},
		cb:     cb,
	}
}

func (p *ReverseProxy) ServeHTTP(c *gin.Context) {
	_, err := p.cb.Execute(func() (interface{}, error) {
		proxy := httputil.NewSingleHostReverseProxy(p.target)
		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = p.target.Scheme
			req.URL.Host = p.target.Host
			req.URL.Path = c.Request.URL.Path
			req.Header = c.Request.Header
		}
		proxy.Transport = p.client.Transport
		proxy.ServeHTTP(c.Writer, c.Request)
		return nil, nil
	})

	if err != nil {
		c.JSON(503, gin.H{"error": "Serviço indisponível"})
	}
}

func (p *ReverseProxy) SetTimeout(timeout time.Duration) {
	p.client.Timeout = timeout
}
