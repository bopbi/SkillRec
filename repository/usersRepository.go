package repository

import (
	"database/sql"

	"github.com/bopbi/SkillRec/checker"
	"github.com/bopbi/SkillRec/entity"
)

// GetUsers will return array of UserResponse
func GetUsers(db *sql.DB) []*entity.UserResponse {

	var userListResponse []*entity.UserResponse
	rows, err := db.Query("SELECT id, name, email FROM users")
	checker.CheckErr(err)
	for rows.Next() {
		userResponse := new(entity.UserResponse)
		err = rows.Scan(&userResponse.ID, &userResponse.Name, &userResponse.Email)
		userListResponse = append(userListResponse, userResponse)
	}
	return userListResponse
}

// GetUsersByID get user by its id
func GetUsersByID(db *sql.DB, id int) *entity.UserResponse {
	userResponse := new(entity.UserResponse)
	rows, err := db.Query("SELECT id, name, email FROM users WHERE id = $1 LIMIT 1", id)
	checker.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(&userResponse.ID, &userResponse.Name, &userResponse.Email)
	}
	return userResponse
}
