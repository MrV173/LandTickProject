package routes

import "github.com/labstack/echo/v4"

func Routeinit(e *echo.Group) {
	UserRoutes(e)
	ProfileRoutes(e)
	StationRoutes(e)
	TicketRoutes(e)
	AuthRoutes(e)
	TransactionRoute(e)
	DestinationRoutes(e)
}
