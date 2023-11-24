package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dakdikduk/galactic-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	spacecraftHandler "github.com/dakdikduk/galactic-api/spacecraft/delivery/http"
	spacecraftRepo "github.com/dakdikduk/galactic-api/spacecraft/repository/mysql"
	spacecraftUc "github.com/dakdikduk/galactic-api/spacecraft/usecase"
)

func main() {
	cfg := config.Get()

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.MysqlDbUser, cfg.MysqlDbPass, cfg.MysqlDbHost, cfg.MysqlDbPort, cfg.MySqlDbName)
	dbConn, err := gorm.Open(mysql.Open(connection), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal(err)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	spacecraftRepo := spacecraftRepo.NewMySQLSpacecraftRepository(dbConn)
	spacecraftUseCase := spacecraftUc.NewSpacecraftUseCase(spacecraftRepo)
	spacecraftHandler.NewSpacecraftHandler(e, spacecraftUseCase)
	// Routes
	e.GET("/_health", healthCheck)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.RESTPort)))
}

// health check
func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
