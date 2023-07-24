package routes

import (
	"landtick/handlers"
	"landtick/pkg/middleware"
	"landtick/pkg/mysql"
	"landtick/repository"

	"github.com/labstack/echo/v4"
)

func TransactionRoute(e *echo.Group) {
	r := repository.TransactionRepository(mysql.DB)
	h := handlers.TransactionHandler(r)

	e.GET("/transactions", (h.FindTransaction))
	e.GET("/transaction/:id", (h.GetTransaction))
	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	e.DELETE("/transaction-delete/:id", (h.DeleteTransaction))
	e.GET("/transaction-user", middleware.Auth(h.GetTransactionUser))
	e.POST("/notification", h.Notification)
}
