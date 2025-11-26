package http

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRoomHandler(srv ports.GameService) gin.HandlerFunc {
	// CreateRoomHandler godoc
	// @Summary Create a game room
	// @Description Creates a new game room for the authenticated user.
	// @Tags Game
	// @Accept json
	// @Produce json
	// @Security BearerAuth
	// @Param request body domain.CreateRoomRequest true "Room payload"
	// @Success 200 {object} domain.GameRoom
	// @Failure 400 {object} map[string]string
	// @Failure 401 {object} map[string]string
	// @Router /game/rooms [post]
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
	// ListRoomsHandler godoc
	// @Summary List game rooms
	// @Description Lists available game rooms.
	// @Tags Game
	// @Produce json
	// @Security BearerAuth
	// @Success 200 {array} domain.GameRoom
	// @Failure 401 {object} map[string]string
	// @Failure 500 {object} map[string]string
	// @Router /game/rooms [get]
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
	// JoinRoomHandler godoc
	// @Summary Join a game room
	// @Description Adds the authenticated user to a room by ID.
	// @Tags Game
	// @Produce json
	// @Security BearerAuth
	// @Param id path int true "Room ID"
	// @Success 200 {object} map[string]string
	// @Failure 400 {object} map[string]string
	// @Failure 401 {object} map[string]string
	// @Router /game/rooms/{id}/join [post]
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
	// LeaveRoomHandler godoc
	// @Summary Leave a game room
	// @Description Removes the authenticated user from a room by ID.
	// @Tags Game
	// @Produce json
	// @Security BearerAuth
	// @Param id path int true "Room ID"
	// @Success 200 {object} map[string]string
	// @Failure 400 {object} map[string]string
	// @Failure 401 {object} map[string]string
	// @Router /game/rooms/{id}/leave [post]
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
	// StartGameHandler godoc
	// @Summary Start a game
	// @Description Starts the game for a given room ID.
	// @Tags Game
	// @Produce json
	// @Security BearerAuth
	// @Param id path int true "Room ID"
	// @Success 200 {object} map[string]string
	// @Failure 400 {object} map[string]string
	// @Failure 401 {object} map[string]string
	// @Router /game/rooms/{id}/start [post]
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
	// AdvancePhaseHandler godoc
	// @Summary Advance game phase
	// @Description Advances the game phase for a room.
	// @Tags Game
	// @Produce json
	// @Security BearerAuth
	// @Param id path int true "Room ID"
	// @Success 200 {object} domain.GameRoom
	// @Failure 400 {object} map[string]string
	// @Failure 401 {object} map[string]string
	// @Router /game/rooms/{id}/phase [post]
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
	// VoteHandler godoc
	// @Summary Submit a vote
	// @Description Records a vote in an active game.
	// @Tags Game
	// @Accept json
	// @Produce json
	// @Security BearerAuth
	// @Param id path int true "Room ID"
	// @Param request body domain.VoteRequest true "Vote payload"
	// @Success 200 {object} map[string]string
	// @Failure 400 {object} map[string]string
	// @Failure 401 {object} map[string]string
	// @Router /game/rooms/{id}/vote [post]
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
	// AbilityHandler godoc
	// @Summary Use an ability
	// @Description Uses a role ability in an active game.
	// @Tags Game
	// @Accept json
	// @Produce json
	// @Security BearerAuth
	// @Param id path int true "Room ID"
	// @Param request body domain.AbilityRequest true "Ability payload"
	// @Success 200 {object} map[string]string
	// @Failure 400 {object} map[string]string
	// @Failure 401 {object} map[string]string
	// @Router /game/rooms/{id}/ability [post]
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
