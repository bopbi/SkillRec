package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bopbi/SkillRec/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {

	port := os.Getenv("PORT")
	dbURL, dbURLExist := os.LookupEnv("DATABASE_URL")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	connStr := ""

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if dbURLExist {
		connStr = dbURL
	} else {
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		connStr = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
			dbUser, dbPassword, dbName)
		e.Use(middleware.CORS())
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	e.File("/", "public/index.html")
	e.GET("/api/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/users", handlers.GetUsers(db))
	e.GET("/api/users/:id/skills", func(c echo.Context) error {
		return c.String(http.StatusOK, "Skill for User with id")
	})
	e.GET("/api/users/:id", handlers.GetUserByID(db))
	e.GET("/api/skills", handlers.GetSkills(db))
	e.GET("/api/skills/:id/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "Users for Skill with id")
	})
	e.GET("/api/skills/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Skill with id")
	})
	e.Logger.Fatal(e.Start(":" + port))
}
