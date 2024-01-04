package databases

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Mariadb *gorm.DB
)

// Database represents a database connection
type Database struct {
	DB *sql.DB
}

// Table represents a database table
type Table struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

// Column represents a column in a database table
type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func init() {
	var err error

	// Mariadb 연결 설정
	Mariadb, err = gorm.Open(mysql.Open(""), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Mariadb:", err)
	}

	sqlDB, err := Mariadb.DB()
	if err != nil {
		log.Fatal("Failed to get underlying sql.DB:", err)
	}

	// // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// sqlDB.SetMaxIdleConns(config.Config.Database.Maria.MaxIdleConns)
	// // SetMaxOpenConns sets the maximum number of open connections to the database.
	// sqlDB.SetMaxOpenConns(config.Config.Database.Maria.MaxOpenConns)
	// // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// sqlDB.SetConnMaxLifetime(time.Duration(config.Config.Database.Maria.ConnMaxLifetime))

	err = sqlDB.Ping()
	if err != nil {
		//서버 중지됨
		log.Fatal("Failed to ping Mariadb:", err)
	}

}

func NewDatabase(dbInstance *sql.DB) (*Database, error) {
	// db, err := sql.Open(driver, source)
	// if err != nil {
	// 	return nil, err
	// }

	return &Database{
		DB: dbInstance,
	}, nil
}

func (db *Database) Close() error {
	return db.DB.Close()
}

// CreateDatabase creates a new database with the given name
func (db *Database) CreateDatabase(databaseName string) error {
	// Database creation logic
	query := "CREATE DATABASE IF NOT EXISTS '?'"
	_, err := db.DB.Exec(query, databaseName)
	if err != nil {
		return err
	}

	return nil
}

// DeleteDatabase deletes the specified database
func (db *Database) DeleteDatabase(databaseName string) error {
	// Database deletion logic
	query := "DROP DATABASE IF EXISTS '?'"
	_, err := db.DB.Exec(query, databaseName)
	if err != nil {
		return err
	}

	return nil
}

// CheckDatabaseExists checks if the specified database exists
func (db *Database) CheckDatabaseExists(databaseName string) (bool, error) {
	// Database existence check logic
	query := "SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = '?'"
	row := db.DB.QueryRow(query, databaseName)
	var dbName string
	err := row.Scan(&dbName)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

// CreateTable creates a new table in the specified database
func (db *Database) CreateTable(databaseName string, tableData *Table) error {
	var args interface{}
	// Table creation logic
	query := "CREATE TABLE IF NOT EXISTS ?.? ("
	args = databaseName
	args = tableData.Name
	for i, column := range tableData.Columns {
		query += "? ? "
		args = column.Name
		args = column.Type
		if i != len(tableData.Columns)-1 {
			query += ", "
		}
	}
	query += ")"

	_, err := db.DB.Exec(query, args)
	if err != nil {
		return err
	}

	return nil
}

// AlterTable alters an existing table in the specified database
func (db *Database) AlterTable(databaseName, tableName string, tableData *Table) error {
	var args interface{}
	// Table alteration logic
	query := "ALTER TABLE ?.? "
	args = databaseName
	args = tableName

	for i, column := range tableData.Columns {
		query += "ADD COLUMN ? ?"
		args = column.Name
		args = column.Type
		if i != len(tableData.Columns)-1 {
			query += ", "
		}
	}

	_, err := db.DB.Exec(query, args)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTable deletes an existing table from the specified database
func (db *Database) DeleteTable(databaseName, tableName string) error {
	// Table deletion logic
	query := "DROP TABLE IF EXISTS ?.?"
	_, err := db.DB.Exec(query, databaseName, tableName)
	if err != nil {
		return err
	}

	return nil
}

// CheckTableExists checks if the specified table exists in the database
func (db *Database) CheckTableExists(databaseName, tableName string) (bool, error) {
	// Table existence check logic
	query := "SELECT table_name FROM information_schema.tables WHERE table_schema = '?' AND table_name = '?'"
	row := db.DB.QueryRow(query, databaseName, tableName)
	var tblName string
	err := row.Scan(&tblName)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

// GetTableData retrieves the data from an existing table in the specified database
func (db *Database) GetTableData(databaseName, tableName string) ([]map[string]interface{}, error) {
	// Table data retrieval logic
	query := "SELECT * FROM ?.?"
	rows, err := db.DB.Query(query, databaseName, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	data := make([]map[string]interface{}, 0)
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		err = rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}

		entry := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			entry[col] = val
		}
		data = append(data, entry)
	}

	return data, nil
}
