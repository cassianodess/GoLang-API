package main

import (
	"main/handler"
	"main/pkg"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.GET("/students", handler.GetStudents)

	e.GET("/students/:id", handler.GetById)

	e.POST("/students", handler.CreateStudent)

	e.PUT("/students/:id", handler.UpdateStudent)

	e.DELETE("/students/:id", handler.DeleteStudent)

	e.Start(":8080")

	defer pkg.Db.Close()

}
