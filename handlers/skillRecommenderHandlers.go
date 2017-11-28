package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/bopbi/SkillRec/repository"
	"github.com/labstack/echo"
)

// GetSkillRecommenderByUserID return array of recommender (userResponse)
func GetSkillRecommenderByUserID(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		userID, _ := strconv.Atoi(paramID)
		userResponse := repository.GetSkillRecommendersByUserID(db, userID)
		return c.JSON(http.StatusOK, userResponse)
	}
}
