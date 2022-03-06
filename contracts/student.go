package contracts

import "main/models"

type DatabaseStudent interface {
	List() ([]models.Student, error)
	Retrieve(id string) (models.Student, error)
	Create(student *models.Student) error
	Update(id string, student *models.Student) (models.Student, error)
	Delete(id string) error
}
