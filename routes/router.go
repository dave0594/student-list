package routes

import (
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
	v1.GET("/students", nil)
	v1.GET("/student/:id", nil)
	v1.POST("/student", nil)
	v1.PUT("/student/:id", nil)
	v1.DELETE("/student/:id", nil)

	return e
}
