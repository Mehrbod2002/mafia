package http

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetProfileHandler godoc
// @Summary Get user profile
// @Description Retrieves the authenticated user's profile details.
// @Tags User
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Profile
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /user/profile [get]
func GetProfileHandler(srv ports.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		profile, err := srv.GetProfile(userID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, profile)
	}
}

// UpdateProfileHandler godoc
// @Summary Update user profile
// @Description Updates the authenticated user's profile information.
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body domain.UpdateProfileRequest true "Profile update payload"
// @Success 200 {object} domain.Profile
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /user/profile [put]
func UpdateProfileHandler(srv ports.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		var req domain.UpdateProfileRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		profile, err := srv.UpdateProfile(userID, req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, profile)
	}
}

// DashboardHandler godoc
// @Summary Get user dashboard
// @Description Retrieves aggregated dashboard information for the authenticated user.
// @Tags User
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /user/dashboard [get]
func DashboardHandler(srv ports.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		dashboard, err := srv.GetDashboard(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, dashboard)
	}
}

// GetWalletHandler godoc
// @Summary Get wallet details
// @Description Retrieves the authenticated user's wallet balances.
// @Tags User
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Wallet
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /user/wallet [get]
func GetWalletHandler(srv ports.WalletService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		wallet, err := srv.GetWallet(userID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, wallet)
	}
}

// PurchaseHandler godoc
// @Summary Initiate wallet purchase
// @Description Initiates a purchase flow for the authenticated user.
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body domain.PurchaseRequest true "Purchase payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /user/purchase [post]
func PurchaseHandler(srv ports.WalletService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		var req domain.PurchaseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		url, err := srv.InitiatePurchase(userID, req.PlanID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"payment_url": url})
	}
}

func BanUserHandler(srv ports.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		err := srv.BanUser(uint(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "banned"})
	}
}

func SuspendUserHandler(srv ports.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		err := srv.SuspendUser(uint(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "suspended"})
	}
}
