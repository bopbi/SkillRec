package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/bopbi/SkillRec/repository"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Login function
func Login(db *sql.DB) echo.HandlerFunc {

	return func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		password = repository.GetMD5Hash(password)

		userResponse := repository.GetUsersByEmailAndPassword(db, email, password)

		if userResponse.ID > 0 {
			// Create token
			token := jwt.New(jwt.SigningMethodHS256)

			// Set claims
			claims := token.Claims.(jwt.MapClaims)
			claims["userID"] = "Jon Snow"
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

			// Generate encoded token and send it as response.
			t, err := token.SignedString([]byte("secretKey"))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]string{
				"token": t,
			})
		}
		return c.String(http.StatusForbidden, "wrong email / password")
	}
}
