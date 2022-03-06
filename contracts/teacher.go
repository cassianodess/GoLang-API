package contracts

import "main/models"

type DatabaseTeacher interface {
	List() ([]models.Teacher, error)
	Retrieve(id string) (models.Teacher, error)
	Create(teacher *models.Teacher) error
	Update(id string, teacher *models.Teacher) (models.Teacher, error)
	Delete(id string) error
}
