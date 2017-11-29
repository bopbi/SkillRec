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

	// static
	e.Static("/static", "public/static")

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
	e.POST("/api/logout", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/api/login", handlers.Login(db))

	// ===
	e.GET("/api/users", handlers.GetUsers(db))
	e.GET("/api/users/:id/recommenders", handlers.GetSkillRecommenderByUserID(db))
	e.GET("/api/users/:id/skills", handlers.GetSkillsByUserID(db))
	e.GET("/api/users/:id", handlers.GetUserByID(db))

	// ===
	e.GET("/api/skills", handlers.GetSkills(db))
	e.GET("/api/skills/:id/users", handlers.GetUsersBySkillID(db))
	e.GET("/api/skills/:id", handlers.GetSkillByID(db))
	e.Logger.Fatal(e.Start(":" + port))
}
