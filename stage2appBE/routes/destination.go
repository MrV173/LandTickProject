package routes

import (
	"landtick/handlers"
	"landtick/pkg/mysql"
	"landtick/repository"

	"github.com/labstack/echo/v4"
)

func DestinationRoutes(e *echo.Group) {
	r := repository.RepositoryDestination(mysql.DB)
	h := handlers.DestinationHandler(r)

	e.GET("/destinations", h.FindDestination)
	e.GET("/destination/:id", h.GetDestination)
	e.POST("/destination", h.CreateDestination)
	e.DELETE("/destination/:id", h.DeleteDestination)
}
