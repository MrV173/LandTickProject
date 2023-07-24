package repository

import (
	"landtick/models"

	"gorm.io/gorm"
)

type TicketRepository interface {
	FindTicket() ([]models.Ticket, error)
	GetTicket(ID int) (models.Ticket, error)
	UpdateTicket(ticket models.Ticket) (models.Ticket, error)
	CreateTicket(ticket models.Ticket) (models.Ticket, error)
	DeleteTicket(ticket models.Ticket) (models.Ticket, error)
}

func RepositoryTicket(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTicket() ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := r.db.Preload("User").Preload("Station").Preload("Destination").Find(&tickets).Error

	return tickets, err
}

func (r *repository) GetTicket(ID int) (models.Ticket, error) {
	var ticket models.Ticket
	err := r.db.Preload("User").First(&ticket, ID).Error

	return ticket, err
}

func (r *repository) CreateTicket(ticket models.Ticket) (models.Ticket, error) {
	err := r.db.Preload("Station").Preload("Destination").Preload("User").Create(&ticket).Error
	return ticket, err
}

func (r *repository) DeleteTicket(ticket models.Ticket) (models.Ticket, error) {
	err := r.db.Delete(&ticket).Error

	return ticket, err
}

func (r *repository) UpdateTicket(ticket models.Ticket) (models.Ticket, error) {
	err := r.db.Save(&ticket).Error

	return ticket, err
}
