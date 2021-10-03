package models

import (
	"net/http"

	"github.com/dave0594/student-list/database"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type Student struct {
	ID          int      `json:"id"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Email       string   `json:"email"`
	PhoneNumber string   `json:"phone_number"`
	CreatedAt   NullTime `json:"created_at"`
	UpdatedAt   NullTime `json:"updated_at"`
}

func GetStudents() (students []Student, err error) {
	q := `SELECT id, first_name, last_name, email, phone_number, created_at, updated_at
			FROM student
			ORDER BY id ASC;`

	createdAtNull := pq.NullTime{}
	updatedAtNull := pq.NullTime{}

	db := database.GetConnection()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		s := Student{}
		err = rows.Scan(
			&s.ID,
			&s.FirstName,
			&s.LastName,
			&s.Email,
			&s.PhoneNumber,
			&createdAtNull,
			&updatedAtNull,
		)
		if err != nil {
			return
		}

		s.CreatedAt = NullTime{createdAtNull.Time}
		s.UpdatedAt = NullTime{updatedAtNull.Time}

		students = append(students, s)
	}
	return students, nil
}

func GetStudentByID(id int) (student Student, err error) {
	q := `SELECT id, first_name, last_name, email, phone_number, created_at, updated_at
			FROM student
			WHERE id = $1;`

	createdAtNull := pq.NullTime{}
	updatedAtNull := pq.NullTime{}

	db := database.GetConnection()
	defer db.Close()

	row := db.QueryRow(q, id)
	err = row.Scan(
		&student.ID,
		&student.FirstName,
		&student.LastName,
		&student.Email,
		&student.PhoneNumber,
		&createdAtNull,
		&updatedAtNull,
	)
	if err != nil {
		return
	}

	student.CreatedAt = NullTime{createdAtNull.Time}
	student.UpdatedAt = NullTime{updatedAtNull.Time}

	return student, nil
}

func CreateStudent(s Student) error {
	q := `INSERT INTO student(first_name, last_name, email, phone_number, created_at)
			VALUES($1, $2, $3, $4, NOW());`

	db := database.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(
		s.FirstName,
		s.LastName,
		s.Email,
		s.PhoneNumber,
	)
	if err != nil {
		return err
	}

	i, _ := row.RowsAffected()
	if i != 1 {
		return echo.NewHTTPError(http.StatusInternalServerError, Message{Message: "1 affected row expected"})
	}
	return nil
}

func UpdateStudent(s Student) error {
	q := `UPDATE student
			SET first_name = $1, last_name = $2, email = $3, phone_number = $4, updated_at = NOW()
			WHERE id = $5;`

	db := database.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(
		s.FirstName,
		s.LastName,
		s.Email,
		s.PhoneNumber,
		s.ID,
	)
	if err != nil {
		return err
	}

	i, _ := row.RowsAffected()
	if i != 1 {
		return echo.NewHTTPError(http.StatusInternalServerError, Message{Message: "1 affected row expected"})
	}
	return nil
}

func DeleteStudent(id int) error {
	q := `DELETE FROM student WHERE id = $1;`

	db := database.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	i, _ := row.RowsAffected()
	if i != 1 {
		return echo.NewHTTPError(http.StatusInternalServerError, Message{Message: "1 affected row expected"})
	}
	return nil
}
