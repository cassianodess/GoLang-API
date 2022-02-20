package contracts

import "main/models"

type Database interface {
	GetAll() ([]models.Student, error)
	GetById(id string) (models.Student, error)
	Create(student *models.Student) error
	Update(id string, student *models.Student) (models.Student, error)
	Delete(id string) error
}
