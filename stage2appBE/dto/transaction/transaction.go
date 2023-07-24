package transactiondto

import "landtick/models"

type CreateTransactionRequest struct {
	UserID   int    `json:"user_id"`
	TicketID int    `json:"ticket_id"`
	Status   string `json:"status"`
	Price    int    `json:"price"`
}

type TransactionResponse struct {
	UserID   int                   `json:"user_id"`
	User     models.UserResponse   `json:"user"`
	TicketID int                   `json:"ticket_id"`
	Ticket   models.TicketResponse `json:"ticket"`
}
