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

// GetUsersByEmailAndPassword check for user by email and password
func GetUsersByEmailAndPassword(db *sql.DB, email string, password string) *entity.UserResponse {
	userResponse := new(entity.UserResponse)
	rows, err := db.Query("SELECT id, name, email FROM users WHERE email LIKE $1 AND password LIKE $2 LIMIT 1", email, password)
	checker.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(&userResponse.ID, &userResponse.Name, &userResponse.Email)
	}
	return userResponse
}

// GetUsersByEmail check for user by email
func GetUsersByEmail(db *sql.DB, email string) *entity.UserResponse {
	userResponse := new(entity.UserResponse)
	rows, err := db.Query("SELECT id, name, email FROM users WHERE email LIKE $1 LIMIT 1", email)
	checker.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(&userResponse.ID, &userResponse.Name, &userResponse.Email)
	}
	return userResponse
}

// InsertUser will return array of UserResponse
func InsertUser(db *sql.DB, name string, email string, password string) *entity.UserResponse {
	rows, err := db.Query("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", name, email, password)
	var lastID = 0
	for rows.Next() {
		err = rows.Scan(&lastID)
	}
	checker.CheckErr(err)
	userResponse := new(entity.UserResponse)
	userResponse.ID = int(lastID)
	userResponse.Name = name
	userResponse.Email = email
	return userResponse
}
