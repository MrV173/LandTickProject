package routes

import (
	"landtick/handlers"
	"landtick/pkg/mysql"
	"landtick/repository"

	"github.com/labstack/echo/v4"
)

func StationRoutes(e *echo.Group) {
	r := repository.RepositoryStation(mysql.DB)
	h := handlers.StationHandler(r)

	e.GET("/stations", h.FindStation)
	e.GET("/station/:id", h.GetStation)
	e.POST("/station", h.CreateStation)
	e.DELETE("/station/:id", h.DeleteStation)
}
