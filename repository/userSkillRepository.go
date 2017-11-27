package repository

import (
	"database/sql"

	"github.com/bopbi/SkillRec/checker"
	"github.com/bopbi/SkillRec/entity"
)

// GetSkillsByUserID return array of SkillResponse
func GetSkillsByUserID(db *sql.DB, id int) []*entity.SkillResponse {
	queryString := `
	SELECT s.id, s.name FROM user_skills us
	INNER JOIN skills s ON us.skill_id = s.id
	WHERE us.user_id = $1
	`
	var skillListResponse []*entity.SkillResponse
	rows, err := db.Query(queryString, id)
	checker.CheckErr(err)
	for rows.Next() {
		skillResponse := new(entity.SkillResponse)
		err = rows.Scan(&skillResponse.ID, &skillResponse.Name)
		skillListResponse = append(skillListResponse, skillResponse)
	}
	return skillListResponse
}

// GetUsersBySkillID return array of UserResponse
func GetUsersBySkillID(db *sql.DB, id int) []*entity.UserResponse {
	queryString := `
	SELECT u.id, u.name, u.email FROM user_skills us
	INNER JOIN users u ON us.user_id = u.id
	WHERE us.skill_id = $1
	`
	var userListResponse []*entity.UserResponse
	rows, err := db.Query(queryString, id)
	checker.CheckErr(err)
	for rows.Next() {
		userResponse := new(entity.UserResponse)
		err = rows.Scan(&userResponse.ID, &userResponse.Name, &userResponse.Email)
		userListResponse = append(userListResponse, userResponse)
	}
	return userListResponse
}
