package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minias/dbgo/databases"
)

// TableController handles the table-related operations
type TableController struct {
	db *databases.Database
}

// NewTableController creates a new TableController instance
func NewTableController(db *databases.Database) *TableController {
	return &TableController{
		db: db,
	}
}

// CreateTable creates a new table in the specified database
func (ctrl *TableController) CreateTable(c echo.Context) error {
	databaseName := c.Param("databaseName")

	// Check if the specified database exists
	exists, err := ctrl.db.CheckDatabaseExists(databaseName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !exists {
		return c.JSON(http.StatusBadRequest, "Database not found")
	}

	// Parse the request body as JSON
	var tableData databases.Table
	if err := c.Bind(&tableData); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid table data")
	}

	// Create the table
	err = ctrl.db.CreateTable(databaseName, &tableData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Table created successfully")
}

// AlterTable alters an existing table in the specified database
func (ctrl *TableController) AlterTable(c echo.Context) error {
	databaseName := c.Param("databaseName")
	tableName := c.Param("tableName")

	// Check if the specified database exists
	exists, err := ctrl.db.CheckDatabaseExists(databaseName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !exists {
		return c.JSON(http.StatusBadRequest, "Database not found")
	}

	// Parse the request body as JSON
	var tableData databases.Table
	if err := c.Bind(&tableData); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid table data")
	}

	// Check if the specified table exists
	tableExists, err := ctrl.db.CheckTableExists(databaseName, tableName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !tableExists {
		return c.JSON(http.StatusBadRequest, "Table not found")
	}

	// Alter the table
	err = ctrl.db.AlterTable(databaseName, tableName, &tableData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Table altered successfully")
}

// DeleteTable deletes an existing table from the specified database
func (ctrl *TableController) DeleteTable(c echo.Context) error {
	databaseName := c.Param("databaseName")
	tableName := c.Param("tableName")

	// Check if the specified database exists
	exists, err := ctrl.db.CheckDatabaseExists(databaseName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !exists {
		return c.JSON(http.StatusBadRequest, "Database not found")
	}

	// Check if the specified table exists
	tableExists, err := ctrl.db.CheckTableExists(databaseName, tableName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !tableExists {
		return c.JSON(http.StatusBadRequest, "Table not found")
	}

	// Delete the table
	err = ctrl.db.DeleteTable(databaseName, tableName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Table deleted successfully")
}

// GetTableData retrieves the data from an existing table in the specified database
func (ctrl *TableController) GetTableData(c echo.Context) error {
	databaseName := c.Param("databaseName")
	tableName := c.Param("tableName")

	// Check if the specified database exists
	exists, err := ctrl.db.CheckDatabaseExists(databaseName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !exists {
		return c.JSON(http.StatusBadRequest, "Database not found")
	}

	// Check if the specified table exists
	tableExists, err := ctrl.db.CheckTableExists(databaseName, tableName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !tableExists {
		return c.JSON(http.StatusBadRequest, "Table not found")
	}

	// Get the table data
	data, err := ctrl.db.GetTableData(databaseName, tableName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}
