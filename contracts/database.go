package contracts

import "main/models"

type Database interface {
	GetAll() ([]models.Student, error)
	GetById(id string) (models.Student, error)
	Create(student *models.Student) error
	Update(id int, student models.Student) error
	Delete(id int) (int64, error)
}
