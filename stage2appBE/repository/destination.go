package repository

import (
	"landtick/models"

	"gorm.io/gorm"
)

type DestinationRepository interface {
	FindDestination() ([]models.Destination, error)
	GetDestination(ID int) (models.Destination, error)
	CreateDestination(destination models.Destination) (models.Destination, error)
	DeleteDestination(destination models.Destination) (models.Destination, error)
}

func RepositoryDestination(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindDestination() ([]models.Destination, error) {
	var destinations []models.Destination
	err := r.db.Find(&destinations).Error

	return destinations, err
}

func (r *repository) GetDestination(ID int) (models.Destination, error) {
	var destination models.Destination
	err := r.db.First(&destination, ID).Error

	return destination, err
}

func (r *repository) CreateDestination(destination models.Destination) (models.Destination, error) {
	err := r.db.Create(&destination).Error
	return destination, err
}

func (r *repository) DeleteDestination(destination models.Destination) (models.Destination, error) {
	err := r.db.Delete(&destination).Error

	return destination, err
}
