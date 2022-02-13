package repository

import (
	"fmt"
	"main/models"
	"main/pkg"
)

type DataBaseImpl struct{}

func (dataBaseImpl DataBaseImpl) GetAll() ([]models.Student, error) {

	results, err := pkg.Db.Query(`select * from students`)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var students []models.Student

	for results.Next() {
		var student models.Student
		err = results.Scan(&student.Id, &student.Name, &student.Course)

		if err != nil {
			println(err)
		}
		students = append(students, student)
	}

	return students, err

}

func (dataBaseImpl DataBaseImpl) GetById(id string) (models.Student, error) {
	student := models.Student{}

	err := pkg.Db.QueryRow(`select * from students where id = $1`, id).Scan(&student.Id, &student.Name, &student.Course)

	return student, err
}

func (dataBaseImpl DataBaseImpl) Create(student *models.Student) error {

	err := pkg.Db.QueryRow(`insert into students (name, course) values ($1, $2) returning id, name, course`, student.Name, student.Course).Scan(&student.Id, &student.Name, &student.Course)

	return err
}

func (dataBaseImpl DataBaseImpl) Update(id string, student *models.Student) error {

	err := pkg.Db.QueryRow(`update students set name=$1, course=$2 where id=$3 returning id, name, course`, student.Name, student.Course, id).Scan(&student.Id, &student.Name, &student.Course)

	return err
}

func (dataBaseImpl DataBaseImpl) Delete(id string) (int64, error) {
	result, err := pkg.Db.Exec(`delete from students where id = $1`, id)
	updates, _ := result.RowsAffected()

	return updates, err
}
