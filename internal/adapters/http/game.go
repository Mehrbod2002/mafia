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
		if err := srv.JoinRoom(uint(roomID), userID); err != nil {
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
		if err := srv.LeaveRoom(uint(roomID), userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "left"})
	}
}

func StartGameHandler(srv ports.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID, _ := strconv.Atoi(c.Param("id"))
		if err := srv.StartGame(uint(roomID)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "game started"})
	}
}

func AdvancePhaseHandler(srv ports.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID, _ := strconv.Atoi(c.Param("id"))
		room, err := srv.AdvancePhase(uint(roomID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, room)
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
		if err := srv.Vote(uint(roomID), userID, req.TargetID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "vote recorded"})
	}
}

func AbilityHandler(srv ports.GameService) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID, _ := strconv.Atoi(c.Param("id"))
		userID := c.GetUint("user_id")
		var req domain.AbilityRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := srv.UseAbility(uint(roomID), userID, req.Ability, req.TargetID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "ability used"})
	}
}
