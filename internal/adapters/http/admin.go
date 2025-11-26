package http

import (
	"mafia/internal/core/domain"
	"mafia/internal/ports"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AdminRoleRoutes(r *gin.RouterGroup, srv ports.AdminService) {
	r.GET("/abilities", ListAbilitiesHandler(srv))
	r.GET("/roles", ListRolesHandler(srv))
	r.POST("/roles", CreateRoleHandler(srv))
	r.PUT("/roles/:id", UpdateRoleHandler(srv))
	r.DELETE("/roles/:id", DeleteRoleHandler(srv))
}

func AdminRuleRoutes(r *gin.RouterGroup, srv ports.AdminService) {
	r.GET("/rules", ListRulesHandler(srv))
	r.POST("/rules", CreateRuleHandler(srv))
	r.PUT("/rules/:id", UpdateRuleHandler(srv))
	r.DELETE("/rules/:id", DeleteRuleHandler(srv))
}

func AdminScenarioRoutes(r *gin.RouterGroup, srv ports.AdminService) {
	r.GET("/scenarios", ListScenariosHandler(srv))
	r.POST("/scenarios", CreateScenarioHandler(srv))
	r.PUT("/scenarios/:id", UpdateScenarioHandler(srv))
	r.DELETE("/scenarios/:id", DeleteScenarioHandler(srv))
}

// ListAbilitiesHandler godoc
// @Summary List available abilities
// @Description Returns all predefined abilities for roles.
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {array} string
// @Router /admin/abilities [get]
func ListAbilitiesHandler(srv ports.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, srv.ListAbilities())
	}
}

// ListRolesHandler godoc
// @Summary List roles
// @Description Lists all configured roles.
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {array} domain.Role
// @Failure 500 {object} map[string]string
// @Router /admin/roles [get]
func ListRolesHandler(srv ports.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, err := srv.ListRoles()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, roles)
	}
}

// CreateRoleHandler godoc
// @Summary Create a role
// @Description Creates a new role definition.
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body domain.CreateRoleRequest true "Role payload"
// @Success 201 {object} domain.Role
// @Failure 400 {object} map[string]string
// @Router /admin/roles [post]
func CreateRoleHandler(srv ports.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

// UpdateRoleHandler godoc
// @Summary Update a role
// @Description Updates an existing role definition.
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Role ID"
// @Param request body domain.CreateRoleRequest true "Role payload"
// @Success 200 {object} domain.Role
// @Failure 400 {object} map[string]string
// @Router /admin/roles/{id} [put]
func UpdateRoleHandler(srv ports.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

// DeleteRoleHandler godoc
// @Summary Delete a role
// @Description Deletes a role by ID.
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param id path int true "Role ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /admin/roles/{id} [delete]
func DeleteRoleHandler(srv ports.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := srv.DeleteRole(uint(id)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}

func ListRulesHandler(srv ports.AdminService) gin.HandlerFunc {
	// ListRulesHandler godoc
	// @Summary List rules
	// @Description Lists all configured rules.
	// @Tags Admin
	// @Produce json
	// @Security BearerAuth
	// @Success 200 {array} domain.GameRule
	// @Failure 500 {object} map[string]string
	// @Router /admin/rules [get]
	return func(c *gin.Context) {
		rules, err := srv.ListRules()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, rules)
	}
}

// CreateRuleHandler godoc
// @Summary Create a rule
// @Description Creates a new game rule.
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body domain.RuleRequest true "Rule payload"
// @Success 201 {object} domain.GameRule
// @Failure 400 {object} map[string]string
// @Router /admin/rules [post]
func CreateRuleHandler(srv ports.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

// UpdateRuleHandler godoc
// @Summary Update a rule
// @Description Updates an existing game rule by ID.
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Rule ID"
// @Param request body domain.RuleRequest true "Rule payload"
// @Success 200 {object} domain.GameRule
// @Failure 400 {object} map[string]string
// @Router /admin/rules/{id} [put]
func UpdateRuleHandler(srv ports.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

// DeleteRuleHandler godoc
// @Summary Delete a rule
// @Description Deletes a game rule by ID.
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param id path int true "Rule ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /admin/rules/{id} [delete]
func DeleteRuleHandler(srv ports.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := srv.DeleteRule(uint(id)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}

// ListScenariosHandler godoc
// @Summary List scenarios
// @Description Lists all configured scenarios.
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {array} domain.Scenario
// @Failure 500 {object} map[string]string
// @Router /admin/scenarios [get]
func ListScenariosHandler(srv ports.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
		scenarios, err := srv.ListScenarios()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, scenarios)
	}
}

// CreateScenarioHandler godoc
// @Summary Create a scenario
// @Description Creates a new scenario with rules and roles.
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body domain.ScenarioRequest true "Scenario payload"
// @Success 201 {object} domain.Scenario
// @Failure 400 {object} map[string]string
// @Router /admin/scenarios [post]
func CreateScenarioHandler(srv ports.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

// UpdateScenarioHandler godoc
// @Summary Update a scenario
// @Description Updates an existing scenario by ID.
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Scenario ID"
// @Param request body domain.ScenarioRequest true "Scenario payload"
// @Success 200 {object} domain.Scenario
// @Failure 400 {object} map[string]string
// @Router /admin/scenarios/{id} [put]
func UpdateScenarioHandler(srv ports.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

// DeleteScenarioHandler godoc
// @Summary Delete a scenario
// @Description Deletes a scenario by ID.
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param id path int true "Scenario ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /admin/scenarios/{id} [delete]
func DeleteScenarioHandler(srv ports.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := srv.DeleteScenario(uint(id)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}
