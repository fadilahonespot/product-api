package repository

import (
	"database/sql"
	"product-api/model"
)

type CategoryRepositoryInterface interface {
	GetAll() ([]model.Category, error)
	Create(category *model.Category) error
	GetByID(id int) (*model.Category, error)
	Update(category *model.Category) error
	Delete(id int) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepositoryInterface {
	return &categoryRepository{db: db}
}

func (repo *categoryRepository) GetAll() ([]model.Category, error) {
	query := "SELECT id, name, description FROM categories"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]model.Category, 0)
	for rows.Next() {
		var c model.Category
		err := rows.Scan(&c.Id, &c.Name, &c.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (repo *categoryRepository) Create(category *model.Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"
	err := repo.db.QueryRow(query, category.Name, category.Description).Scan(&category.Id)
	return err
}

func (repo *categoryRepository) GetByID(id int) (*model.Category, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"
	var c model.Category
	err := repo.db.QueryRow(query, id).Scan(&c.Id, &c.Name, &c.Description)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (repo *categoryRepository) Update(category *model.Category) error {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"
	_, err := repo.db.Exec(query, category.Name, category.Description, category.Id)
	return err
}

func (repo *categoryRepository) Delete(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := repo.db.Exec(query, id)
	return err
}