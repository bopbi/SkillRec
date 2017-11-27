package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/bopbi/SkillRec/repository"
	"github.com/labstack/echo"
)

// GetSkillsByUserID return skill with id
func GetSkillsByUserID(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		skillID, _ := strconv.Atoi(paramID)
		skillResponse := repository.GetSkillsByUserID(db, skillID)
		return c.JSON(http.StatusOK, skillResponse)
	}
}

// GetUsersByUserID return user by id
func GetUsersBySkillID(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		userID, _ := strconv.Atoi(paramID)
		userResponse := repository.GetUsersBySkillID(db, userID)
		return c.JSON(http.StatusOK, userResponse)
	}
}
