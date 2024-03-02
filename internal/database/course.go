package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{
		db: db,
	}
}

func (c *Course) CreateCourse(name, description, categoryId string) (Course, error) {

	course := Course{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		CategoryID:  categoryId,
	}

	stmt, err := c.db.Prepare(`INSERT INTO courses (id,name,description, category_id) values ($1,$2,$3, $4)`)

	if err != nil {
		return Course{}, err
	}

	_, err = stmt.Exec(course.ID, course.Name, course.Description, course.CategoryID)

	if err != nil {
		return Course{}, err
	}

	return course, nil
}

func (c *Course) ListAll() ([]Course, error) {
	rows, err := c.db.Query(`SELECT id,name,description FROM courses`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	courses := []Course{}

	for rows.Next() {
		var course Course

		if err := rows.Scan(&course.ID, &course.Name, &course.Description); err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func (c *Course) FindByCategoryId(categoryId string) ([]Course, error) {
	fmt.Println("chamou")

	rows, err := c.db.Query(`SELECT * FROM courses WHERE category_id = $1`, categoryId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	courses := []Course{}

	for rows.Next() {

		var course Course

		if err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID); err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}
