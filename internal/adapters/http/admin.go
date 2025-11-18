package http

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AdminRoleRoutes(r *gin.RouterGroup, srv ports.AdminService) {
	r.GET("/roles", func(c *gin.Context) {
		roles, err := srv.ListRoles()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, roles)
	})
	r.POST("/roles", func(c *gin.Context) {
		var req domain.CreateRoleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		role, err := srv.CreateRole(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, role)
	})
	r.PUT("/roles/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var req domain.CreateRoleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		role, err := srv.UpdateRole(uint(id), req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, role)
	})
	r.DELETE("/roles/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := srv.DeleteRole(uint(id)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	})
}

func AdminRuleRoutes(r *gin.RouterGroup, srv ports.AdminService) {
	r.GET("/rules", func(c *gin.Context) {
		rules, err := srv.ListRules()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, rules)
	})
	r.POST("/rules", func(c *gin.Context) {
		var req domain.RuleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		rule, err := srv.CreateRule(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, rule)
	})
	r.PUT("/rules/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var req domain.RuleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		rule, err := srv.UpdateRule(uint(id), req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, rule)
	})
	r.DELETE("/rules/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := srv.DeleteRule(uint(id)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	})
}

func AdminScenarioRoutes(r *gin.RouterGroup, srv ports.AdminService) {
	r.GET("/scenarios", func(c *gin.Context) {
		scenarios, err := srv.ListScenarios()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, scenarios)
	})
	r.POST("/scenarios", func(c *gin.Context) {
		var req domain.ScenarioRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		scenario, err := srv.CreateScenario(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, scenario)
	})
	r.PUT("/scenarios/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var req domain.ScenarioRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		scenario, err := srv.UpdateScenario(uint(id), req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, scenario)
	})
	r.DELETE("/scenarios/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := srv.DeleteScenario(uint(id)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	})
}
