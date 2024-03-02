package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

func (c *Category) CreateCategory(name, description string) (Category, error) {

	cat := Category{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
	}

	stmt, err := c.db.Prepare(`INSERT INTO categories (id,name,description) values ($1,$2,$3)`)

	if err != nil {
		return Category{}, err
	}

	_, err = stmt.Exec(cat.ID, cat.Name, cat.Description)

	if err != nil {
		return Category{}, err
	}

	return cat, nil
}

func (c *Category) FindAll() ([]Category, error) {

	rows, err := c.db.Query(`SELECT * FROM categories`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := []Category{}

	for rows.Next() {

		var category Category

		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (c *Category) FindByCourseId(courseId string) (Category, error) {

	var cat Category

	err := c.db.QueryRow(`
    SELECT c.id, c.name, c.description FROM categories
    c JOIN courses co ON c.id = co.category_id
    WHERE co.id = $1
  `, courseId).Scan(&cat.ID, &cat.Name, &cat.Description)

	if err != nil {
		return Category{}, err
	}

	return cat, nil
}
