package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/_health", healthCheck)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
