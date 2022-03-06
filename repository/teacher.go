package repository

import (
	"fmt"
	"main/contracts"
	"main/models"
	"main/pkg"
)

type DatabaseTeacherImpl struct{}

var dataBase contracts.DatabaseTeacher = new(DatabaseTeacherImpl)

func (dataBaseImpl DatabaseTeacherImpl) List() ([]models.Teacher, error) {
	var teachers []models.Teacher

	result := pkg.Db.Order("id").Find(&teachers)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return teachers, result.Error

}

func (dataBaseImpl DatabaseTeacherImpl) Retrieve(id string) (models.Teacher, error) {
	var teacher models.Teacher

	result := pkg.Db.First(&teacher, id)

	return teacher, result.Error
}

func (dataBaseImpl DatabaseTeacherImpl) Create(teacher *models.Teacher) error {

	result := pkg.Db.Create(&teacher)

	return result.Error
}

func (dataBaseImpl DatabaseTeacherImpl) Update(id string, teacher *models.Teacher) (models.Teacher, error) {
	var oldTeacher models.Teacher

	result := pkg.Db.First(&oldTeacher, id)

	oldTeacher.Name = teacher.Name
	oldTeacher.Subject = teacher.Subject

	pkg.Db.Save(&oldTeacher)

	return oldTeacher, result.Error
}

func (dataBaseImpl DatabaseTeacherImpl) Delete(id string) error {

	_, err := dataBase.Retrieve(id)

	if err != nil {
		return err
	}

	result := pkg.Db.Delete(&models.Teacher{}, id)

	return result.Error
}
