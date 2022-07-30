package main

import (
	"go-was-example/api-db-server/database"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// DB Instance
	rdbHandler := database.RDBHandler{
		UserName:      "dev",
		Password:      "jeongchul!@#",
		ServerAddress: "[GOOGLE_CLOUD_SQL_INTERNAL_IP]",
		DbName:        "dev",
	}

	rdbHandler.Connect()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
