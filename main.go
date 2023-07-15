package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/minias/dbgo/config"
	"github.com/minias/dbgo/controller" // 변경된 import 경로
	"github.com/minias/dbgo/database"
)

var (
	cfg   *config.Config
	db    *database.Database
	users map[string]string
)

func init() {
	var err error
	// Load configuration from config.yml
	cfg, err = config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Configure middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize database connection
	// db, err := sql.open
	//  .Connect.ConnectDB(cfg.Database.DBDriver, cfg.Database.DBSource)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }
	// defer db.Close()

	// Create database instance
	dbInstance := database.NewDatabase(db)

	// Create controllers
	dbController := controller.NewDatabaseController(dbInstance)
	tableController := controller.NewTableController(dbInstance)

	// Middleware
	//e.Use(middleware.JWT([]byte(cfg.JWT.Secret)))

	// Group routes
	api := e.Group("/api")

	// Database routes
	databaseGroup := api.Group("/database")
	databaseGroup.POST("", dbController.CreateDatabase)
	databaseGroup.DELETE("/:databaseName", dbController.DeleteDatabase)

	// Table routes
	tableGroup := api.Group("/table/:databaseName")
	tableGroup.POST("", tableController.CreateTable)
	tableGroup.PUT("/:tableName", tableController.AlterTable)
	tableGroup.DELETE("/:tableName", tableController.DeleteTable)
	tableGroup.GET("/:tableName", tableController.GetTableData)

	// User routes
	users := map[string]string{}
	userController := controller.NewUserController(users, cfg.JWT.Secret, time.Duration(cfg.JWT.ExpireTime))
	userGroup := api.Group("/user")
	userGroup.POST("/signup", userController.Signup)
	userGroup.POST("/login", userController.Login)
	userGroup.GET("/signout", userController.Signout)

	// Start the server
	if cfg.Server.Protocol == "https" {
		e.Pre(middleware.HTTPSRedirect())
		err = e.StartTLS(cfg.Server.Port, cfg.Server.CertFile, cfg.Server.KeyFile)
	} else {
		err = e.Start(cfg.Server.Port)
	}
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
