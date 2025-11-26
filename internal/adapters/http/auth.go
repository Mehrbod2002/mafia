package http

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHandler godoc
// @Summary Request OTP for registration
// @Description Sends a one-time password to start a new account registration.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body domain.RegisterRequest true "Registration payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /auth/register [post]
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

// VerifyOTPHandler godoc
// @Summary Verify an OTP code and issue a token
// @Description Confirms the OTP sent to the user and returns an access token.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body domain.VerifyOTPRequest true "OTP verification payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/verify [post]
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

// LoginHandler godoc
// @Summary Request OTP for login
// @Description Sends a one-time password to authenticate an existing user.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body domain.LoginRequest true "Login payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /auth/login [post]
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
