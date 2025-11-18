package http

import (
	"mafia/internal/ports"
	"net/http"
	"github.com/gin-gonic/gin"
)

func ${handler}Handler(s ports.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "implemented"})
	}
}
