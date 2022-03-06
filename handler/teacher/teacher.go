package teacher

import (
	"main/models"
	repo "main/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var databaseTeacher repo.DatabaseTeacherImpl

func List(c echo.Context) error {

	teachers, err := databaseTeacher.List()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Try again")
	}

	return c.JSON(http.StatusOK, teachers)
}

func Retrieve(c echo.Context) error {
	id := c.Param("id")

	teacher, err := databaseTeacher.Retrieve(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, "id not found")
	}

	return c.JSON(http.StatusOK, teacher)

}

func Create(c echo.Context) error {
	teacher := models.Teacher{}
	err := c.Bind(&teacher)
	if err != nil {
		return err
	}

	if len(teacher.Name) > 0 && len(teacher.Subject) > 0 {

		err := databaseTeacher.Create(&teacher)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusCreated, teacher)
	}
	return c.JSON(http.StatusBadRequest, "invalid fields")
}

func Update(c echo.Context) error {

	newTeacher := models.Teacher{}

	err := c.Bind(&newTeacher)

	if err != nil {
		return err
	}

	id := c.Param("id")

	teacher, err := databaseTeacher.Update(id, &newTeacher)

	if err != nil {
		return c.JSON(http.StatusNotFound, "id not found")

	}
	return c.JSON(http.StatusOK, teacher)

}

func Delete(c echo.Context) error {

	id := c.Param("id")

	err := databaseTeacher.Delete(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, "id not found")
	}

	return c.JSON(http.StatusOK, "teacher deleted successful")

}
