package routes

import (
	"github.com/dave0594/student-list/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	// Routes

	v1 := e.Group("/v1")

	// Students
	v1.GET("/students", api.GetStudents)
	v1.GET("/students/:id", api.GetStudentByID)
	v1.POST("/students", api.CreateStudent)
	v1.PUT("/students/:id", api.UpdateStudent)
	v1.DELETE("/students/:id", api.DeleteStudent)

	return e
}
