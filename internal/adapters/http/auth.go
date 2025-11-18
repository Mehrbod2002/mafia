package http

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(srv ports.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := srv.Register(req.Phone)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "otp sent"})
	}
}

func VerifyOTPHandler(srv ports.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.VerifyOTPRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token, userID, err := srv.VerifyOTP(req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token, "user_id": userID})
	}
}

func LoginHandler(srv ports.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := srv.Login(req.Phone)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "otp sent"})
	}
}

func AuthMiddleware(srv ports.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}
		token = token[len("Bearer "):]
		userID, err := srv.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}

func AdminMiddleware(srv ports.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		isAdmin, err := srv.IsAdmin(userID)
		if err != nil || !isAdmin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin only"})
			return
		}
		c.Next()
	}
}
