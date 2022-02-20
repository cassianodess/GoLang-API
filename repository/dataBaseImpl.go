package repository

import (
	"fmt"
	"main/contracts"
	"main/models"
	"main/pkg"
)

type DataBaseImpl struct{}

var dataBase contracts.Database = new(DataBaseImpl)

func (dataBaseImpl DataBaseImpl) GetAll() ([]models.Student, error) {
	var students []models.Student

	result := pkg.Db.Order("id").Find(&students)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return students, result.Error

}

func (dataBaseImpl DataBaseImpl) GetById(id string) (models.Student, error) {
	var student models.Student

	result := pkg.Db.First(&student, id)

	return student, result.Error
}

func (dataBaseImpl DataBaseImpl) Create(student *models.Student) error {

	result := pkg.Db.Create(&student)

	return result.Error
}

func (dataBaseImpl DataBaseImpl) Update(id string, student *models.Student) (models.Student, error) {
	var oldStudent models.Student

	result := pkg.Db.First(&oldStudent, id)

	oldStudent.Name = student.Name
	oldStudent.Course = student.Course

	pkg.Db.Save(&oldStudent)

	return oldStudent, result.Error
}

func (dataBaseImpl DataBaseImpl) Delete(id string) error {

	_, err := dataBase.GetById(id)

	if err != nil {
		return err
	}

	result := pkg.Db.Delete(&models.Student{}, id)

	return result.Error
}
