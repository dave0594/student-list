package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/dave0594/student-list/models"
	"github.com/labstack/echo/v4"
)

func GetStudents(c echo.Context) error {
	students, err := models.GetStudents()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.Message{Message: err.Error()})
	}
	if students == nil {
		return c.JSON(http.StatusNotFound, []models.Student{})
	}
	return c.JSON(http.StatusOK, students)
}

func GetStudentByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	student, err := models.GetStudentByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, nil)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, student)
}

func CreateStudent(c echo.Context) error {
	student := models.Student{}

	err := c.Bind(&student)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Message{Message: err.Error()})
	}

	err = models.CreateStudent(student)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, models.Message{Message: "student created successfully"})
}

func UpdateStudent(c echo.Context) error {
	student := models.Student{}

	err := c.Bind(&student)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = models.UpdateStudent(student)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.Message{Message: "role updated successfully"})
}

func DeleteStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := models.DeleteStudent(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.Message{Message: "role deleted successfully"})
}
