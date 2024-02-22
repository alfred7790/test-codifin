package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MiddlewareService interface {
	AddCORS() gin.HandlerFunc
}

type MiddlewareServiceImpl struct{}

func NewMiddlewareService() *MiddlewareServiceImpl {
	return &MiddlewareServiceImpl{}
}

func (m *MiddlewareServiceImpl) AddCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, Content-Length, Origin")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // Seconds

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
