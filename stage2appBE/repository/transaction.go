package repository

import (
	"landtick/models"

	"gorm.io/gorm"
)

type RepositoryTransaction interface {
	FindTransaction() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	GetTransactionUser(UserID int) ([]models.Transaction, error)
	UpdateTransaction(status string, orderId int) (models.Transaction, error)
}

func TransactionRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Ticket").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Ticket").Preload("User").Create(&transaction).Error
	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&transaction, ID).Scan(&transaction).Error
	return transaction, err
}

func (r *repository) GetTransactionUser(UserID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Ticket").Where("user_id = ?", UserID).Find(&transactions).Error

	return transactions, err
}

func (r *repository) UpdateTransaction(status string, orderId int) (models.Transaction, error) {
	var transaction models.Transaction

	r.db.First(&transaction, orderId)

	if status != transaction.Status && status == "success" {
		var ticket models.Ticket
		r.db.First(&ticket, transaction.Ticket.ID)
	}

	transaction.Status = status
	err := r.db.Save(&transaction).Error

	return transaction, err
}
