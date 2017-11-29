package repository

import (
	"database/sql"

	"github.com/bopbi/SkillRec/checker"
	"github.com/bopbi/SkillRec/entity"
)

// GetSkills return array of SkillResponse
func GetSkills(db *sql.DB) []*entity.SkillResponse {
	var skillListResponse []*entity.SkillResponse
	rows, err := db.Query("SELECT id, name FROM skills")
	checker.CheckErr(err)
	for rows.Next() {
		skillResponse := new(entity.SkillResponse)
		err = rows.Scan(&skillResponse.ID, &skillResponse.Name)
		skillListResponse = append(skillListResponse, skillResponse)
	}
	return skillListResponse
}

// GetSkillByID get user by its id
func GetSkillByID(db *sql.DB, id int) *entity.SkillResponse {
	skillResponse := new(entity.SkillResponse)
	rows, err := db.Query("SELECT id, name FROM skills WHERE id = $1 LIMIT 1", id)
	checker.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(&skillResponse.ID, &skillResponse.Name)
	}
	return skillResponse
}

// GetSkillByName get user by its id
func GetSkillByName(db *sql.DB, name string) *entity.SkillResponse {
	skillResponse := new(entity.SkillResponse)
	rows, err := db.Query("SELECT id, name FROM skills WHERE name LIKE $1 LIMIT 1", name)
	checker.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(&skillResponse.ID, &skillResponse.Name)
	}
	return skillResponse
}
