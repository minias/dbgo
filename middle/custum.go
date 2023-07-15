package middle

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePasswords compares the provided password with the hashed password
func ComparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GetSigningMethod returns the JWT signing method
func GetSigningMethod() jwt.SigningMethod {
	return jwt.SigningMethodHS256
}

// LogMiddleware logs the incoming requests
func LogMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			if err != nil {
				c.Error(err)
			}
			stop := time.Now()

			path := c.Path()
			method := c.Request().Method
			latency := stop.Sub(start).String()

			message := strings.Join([]string{method, path, latency}, " ")

			c.Logger().Info(message)

			return nil
		}
	}
}
