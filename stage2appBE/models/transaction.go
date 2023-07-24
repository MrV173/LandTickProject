package models

type Transaction struct {
	ID       int            `json:"id" gorm:"primaryKey:autoIncrement"`
	UserID   int            `json:"user_id"`
	User     UserResponse   `json:"user"`
	TicketID int            `json:"ticket_id"`
	Ticket   TicketResponse `json:"ticket"`
	Status   string         `json:"status" gorm:"type:varchar(255)"`
	Price    int            `json:"price"`
}

type TransactionResponse struct {
	ID       int            `json:"id"`
	UserID   int            `json:"user_id"`
	User     UserResponse   `json:"user"`
	TicketID int            `json:"ticket_id"`
	Ticket   TicketResponse `json:"ticket"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
