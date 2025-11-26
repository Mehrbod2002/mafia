package http

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListShopItemsHandler(srv ports.ShopService) gin.HandlerFunc {
	// ListShopItemsHandler godoc
	// @Summary List shop items
	// @Description Returns all shop items available for purchase.
	// @Tags Shop
	// @Produce json
	// @Security BearerAuth
	// @Success 200 {array} domain.ShopItem
	// @Failure 401 {object} map[string]string
	// @Failure 500 {object} map[string]string
	// @Router /shop/items [get]
	return func(c *gin.Context) {
		items, err := srv.ListItems()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func PurchaseItemHandler(srv ports.ShopService) gin.HandlerFunc {
	// PurchaseItemHandler godoc
	// @Summary Purchase a shop item
	// @Description Purchases an item from the shop for the authenticated user.
	// @Tags Shop
	// @Accept json
	// @Produce json
	// @Security BearerAuth
	// @Param request body domain.PurchaseItemRequest true "Purchase payload"
	// @Success 200 {object} domain.ShopItem
	// @Failure 400 {object} map[string]string
	// @Failure 401 {object} map[string]string
	// @Router /shop/purchase [post]
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		var req domain.PurchaseItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		item, err := srv.PurchaseItem(userID, req.ItemID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, item)
	}
}

func CreateShopItemHandler(srv ports.ShopService) gin.HandlerFunc {
	// CreateShopItemHandler godoc
	// @Summary Create a shop item
	// @Description Adds a new item to the shop catalog.
	// @Tags Admin
	// @Accept json
	// @Produce json
	// @Security BearerAuth
	// @Param request body domain.ShopItem true "Shop item payload"
	// @Success 201 {object} domain.ShopItem
	// @Failure 400 {object} map[string]string
	// @Failure 401 {object} map[string]string
	// @Router /admin/shop/items [post]
	return func(c *gin.Context) {
		var item domain.ShopItem
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		created, err := srv.CreateItem(item)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, created)
	}
}

func UpdateShopItemHandler(srv ports.ShopService) gin.HandlerFunc {
	// UpdateShopItemHandler godoc
	// @Summary Update a shop item
	// @Description Updates an existing shop item by ID.
	// @Tags Admin
	// @Accept json
	// @Produce json
	// @Security BearerAuth
	// @Param id path int true "Item ID"
	// @Param request body domain.ShopItem true "Shop item payload"
	// @Success 200 {object} domain.ShopItem
	// @Failure 400 {object} map[string]string
	// @Failure 401 {object} map[string]string
	// @Router /admin/shop/items/{id} [put]
	return func(c *gin.Context) {
		var item domain.ShopItem
		id, _ := strconv.Atoi(c.Param("id"))
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		item.ID = uint(id)
		updated, err := srv.UpdateItem(item)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, updated)
	}
}

func DeleteShopItemHandler(srv ports.ShopService) gin.HandlerFunc {
	// DeleteShopItemHandler godoc
	// @Summary Delete a shop item
	// @Description Deletes an existing shop item by ID.
	// @Tags Admin
	// @Produce json
	// @Security BearerAuth
	// @Param id path int true "Item ID"
	// @Success 200 {object} map[string]string
	// @Failure 400 {object} map[string]string
	// @Failure 401 {object} map[string]string
	// @Router /admin/shop/items/{id} [delete]
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := srv.DeleteItem(uint(id)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}
