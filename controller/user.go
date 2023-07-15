package controller

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/minias/dbgo/middle"
)

// UserController handles the user-related operations
type UserController struct {
	users         map[string]string
	jwtSecret     string
	jwtExpireTime time.Duration
}

// NewUserController creates a new UserController instance
func NewUserController(users map[string]string, jwtSecret string, jwtExpireTime time.Duration) *UserController {
	return &UserController{
		users:         users,
		jwtSecret:     jwtSecret,
		jwtExpireTime: jwtExpireTime,
	}
}

// Signup creates a new user
func (ctrl *UserController) Signup(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Check if user already exists
	if _, ok := ctrl.users[email]; ok {
		return c.JSON(http.StatusBadRequest, "User already exists")
	}

	// Hash the password
	hashedPassword, err := middle.HashPassword(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Store the user
	ctrl.users[email] = hashedPassword

	return c.JSON(http.StatusOK, "User created successfully")
}

// Signout signs out the user by revoking the JWT token
func (ctrl *UserController) Signout(c echo.Context) error {
	// Revoke the JWT token by setting an empty token
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "User signed out successfully")
}

// Login authenticates the user and returns a JWT token
func (ctrl *UserController) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Check if user exists
	hashedPassword, ok := ctrl.users[email]
	if !ok {
		return c.JSON(http.StatusUnauthorized, "Invalid email or password")
	}

	// Compare the passwords
	err := middle.ComparePasswords(hashedPassword, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid email or password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(ctrl.jwtExpireTime).Unix(),
	})

	signedToken, err := token.SignedString([]byte(ctrl.jwtSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Set the JWT token as a cookie
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Path:     "/",
		Expires:  time.Now().Add(ctrl.jwtExpireTime),
		HttpOnly: true,
		Secure:   true,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "User logged in successfully")
}
