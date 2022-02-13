package handler

import (
	"main/models"
	repo "main/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var database repo.DataBaseImpl

func GetStudents(c echo.Context) error {

	students, err := database.GetAll()
	if err != nil {

		return c.JSON(http.StatusBadRequest, "Try again")
	}

	return c.JSON(http.StatusOK, students)
}

func GetById(c echo.Context) error {
	id := c.Param("id")

	student, err := database.GetById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, "id not found")
	}

	return c.JSON(http.StatusOK, student)

}

func CreateStudent(c echo.Context) error {
	student := models.Student{}
	err := c.Bind(&student)
	if err != nil {
		return err
	}

	if len(student.Name) > 0 && len(student.Course) > 0 {

		err := database.Create(&student)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusCreated, student)
	}
	return c.JSON(http.StatusBadRequest, "invalid fields")

}

func UpdateStudent(c echo.Context) error {

	newStudent := models.Student{}

	err := c.Bind(&newStudent)

	if err != nil {
		return err
	}

	id := c.Param("id")

	err = database.Update(id, &newStudent)

	if err != nil {
		return c.JSON(http.StatusNotFound, "id not found")

	}
	return c.JSON(http.StatusOK, newStudent)

}

func DeleteStudent(c echo.Context) error {

	id := c.Param("id")

	updates, err := database.Delete(id)

	if err != nil || updates < 1 {
		return c.JSON(http.StatusNotFound, "id not found")
	}

	return c.JSON(http.StatusOK, "student deleted successful")

}
