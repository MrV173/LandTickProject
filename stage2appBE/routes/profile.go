package routes

import (
	"landtick/handlers"
	"landtick/pkg/mysql"
	"landtick/repository"

	"github.com/labstack/echo/v4"
)

func ProfileRoutes(e *echo.Group) {
	r := repository.RepositoryProfile(mysql.DB)
	h := handlers.ProfileHandler(r)

	e.GET("/profiles", h.FindProfile)
	e.GET("/profile/:id", h.GetProfile)
	e.POST("/profile", h.CreateProfile)
	e.PATCH("/profile/:id", h.UpdateProfile)
	e.DELETE("/profile/:id", h.DeleteProfile)
}
