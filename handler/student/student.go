package student

import (
	"main/models"
	repo "main/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var dataBaseStudent repo.DataBaseImpl

func List(c echo.Context) error {

	students, err := dataBaseStudent.List()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Try again")
	}

	return c.JSON(http.StatusOK, students)
}

func Retrieve(c echo.Context) error {
	id := c.Param("id")

	student, err := dataBaseStudent.Retrieve(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, "id not found")
	}

	return c.JSON(http.StatusOK, student)

}

func Create(c echo.Context) error {
	student := models.Student{}
	err := c.Bind(&student)
	if err != nil {
		return err
	}

	if len(student.Name) > 0 && len(student.Course) > 0 {

		err := dataBaseStudent.Create(&student)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusCreated, student)
	}
	return c.JSON(http.StatusBadRequest, "invalid fields")
}

func Update(c echo.Context) error {

	newStudent := models.Student{}

	err := c.Bind(&newStudent)

	if err != nil {
		return err
	}

	id := c.Param("id")

	student, err := dataBaseStudent.Update(id, &newStudent)

	if err != nil {
		return c.JSON(http.StatusNotFound, "id not found")

	}
	return c.JSON(http.StatusOK, student)

}

func Delete(c echo.Context) error {

	id := c.Param("id")

	err := dataBaseStudent.Delete(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, "id not found")
	}

	return c.JSON(http.StatusOK, "student deleted successful")

}
