package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minias/dbgo/databases"
)

// DatabaseController handles the database-related operations
type DatabaseController struct {
	db *databases.Database
}

// NewDatabaseController creates a new DatabaseController instance
func NewDatabaseController(db *databases.Database) *DatabaseController {
	return &DatabaseController{
		db: db,
	}
}

// CreateDatabase creates a new database
func (ctrl *DatabaseController) CreateDatabase(c echo.Context) error {
	databaseName := c.FormValue("databaseName")

	err := ctrl.db.CreateDatabase(databaseName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Database created successfully")
}

// DeleteDatabase deletes an existing database
func (ctrl *DatabaseController) DeleteDatabase(c echo.Context) error {
	databaseName := c.FormValue("databaseName")

	exists, err := ctrl.db.CheckDatabaseExists(databaseName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if !exists {
		return c.JSON(http.StatusBadRequest, "Database not found")
	}

	err = ctrl.db.DeleteDatabase(databaseName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Database deleted successfully")
}
