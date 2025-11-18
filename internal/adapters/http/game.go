package http

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRoomHandler(srv ports.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		var req domain.CreateRoomRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		room, err := srv.CreateRoom(userID, req.Type)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, room)
	}
}

func ListRoomsHandler(srv ports.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		rooms, err := srv.ListRooms()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, rooms)
	}
}

func JoinRoomHandler(srv ports.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID, _ := strconv.Atoi(c.Param("id"))
		userID := c.GetUint("user_id")
		err := srv.JoinRoom(uint(roomID), userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "joined"})
	}
}

func LeaveRoomHandler(srv ports.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID, _ := strconv.Atoi(c.Param("id"))
		userID := c.GetUint("user_id")
		err := srv.LeaveRoom(uint(roomID), userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "left"})
	}
}

func StartGameHandler(srv ports.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID, _ := strconv.Atoi(c.Param("id"))
		err := srv.StartGame(uint(roomID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "started"})
	}
}

func VoteHandler(srv ports.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID, _ := strconv.Atoi(c.Param("id"))
		userID := c.GetUint("user_id")
		var req domain.VoteRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := srv.Vote(uint(roomID), userID, req.TargetID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "voted"})
	}
}

func UseAbilityHandler(srv ports.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID, _ := strconv.Atoi(c.Param("id"))
		userID := c.GetUint("user_id")
		var req domain.AbilityRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := srv.UseAbility(uint(roomID), userID, req.Ability, req.TargetID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "ability used"})
	}
}
