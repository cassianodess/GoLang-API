package main

import (
	"main/handler/student"
	"main/handler/teacher"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-Session-Id"},
	}))

	group := e.Group("/api")

	group.GET("/students", student.List)

	group.GET("/students/:id", student.Retrieve)

	group.POST("/students", student.Create)

	group.PUT("/students/:id", student.Update)

	group.DELETE("/students/:id", student.Delete)

	group.GET("/teachers", teacher.List)

	group.GET("/teachers/:id", teacher.Retrieve)

	group.POST("/teachers", teacher.Create)

	group.PUT("/teachers/:id", teacher.Update)

	group.DELETE("/teachers/:id", teacher.Delete)

	e.Start(":8080")

}
