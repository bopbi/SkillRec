package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/bopbi/SkillRec/repository"

	"github.com/labstack/echo"
)

// GetUsers return all users
func GetUsers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		usersResponse := repository.GetUsers(db)
		return c.JSON(http.StatusOK, usersResponse)
	}
}

// GetUserByID return user by id
func GetUserByID(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		userID, _ := strconv.Atoi(paramID)
		userResponse := repository.GetUsersByID(db, userID)
		return c.JSON(http.StatusOK, userResponse)
	}
}
