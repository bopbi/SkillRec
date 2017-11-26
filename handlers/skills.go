package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/bopbi/SkillRec/repository"

	"github.com/labstack/echo"
)

// GetSkills return all skills
func GetSkills(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		skillsResponse := repository.GetSkills(db)
		return c.JSON(http.StatusOK, skillsResponse)
	}
}

// GetSkillByID return skill with id
func GetSkillByID(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		skillID, _ := strconv.Atoi(paramID)
		skillResponse := repository.GetSkillByID(db, skillID)
		return c.JSON(http.StatusOK, skillResponse)
	}
}
