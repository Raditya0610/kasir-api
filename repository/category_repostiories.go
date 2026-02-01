package repository

import (
	"database/sql"
	"errors"
	"kasir-api/models"
	"time"
)

type CategoryRepository interface {
	FetchAll() ([]models.Category, error)
	FetchByID(id int) (models.Category, error)
	Store(category *models.Category) error
	Update(category *models.Category) error
	Delete(id int) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) FetchAll() ([]models.Category, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM categories WHERE deleted_at IS NULL`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *categoryRepository) FetchByID(id int) (models.Category, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM categories WHERE id = $1 AND deleted_at IS NULL`

	var c models.Category
	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return c, errors.New("category not found")
		}
		return c, err
	}
	return c, nil
}

func (r *categoryRepository) Store(c *models.Category) error {
	query := `
		INSERT INTO categories (name, description, created_at, updated_at) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id
	`
	now := time.Now()
	err := r.db.QueryRow(query, c.Name, c.Description, now, now).Scan(&c.ID)
	if err != nil {
		return err
	}
	c.CreatedAt = now
	c.UpdatedAt = now
	return nil
}

func (r *categoryRepository) Update(c *models.Category) error {
	query := `
		UPDATE categories 
		SET name = $1, description = $2, updated_at = $3
		WHERE id = $4 AND deleted_at IS NULL
	`
	c.UpdatedAt = time.Now()
	res, err := r.db.Exec(query, c.Name, c.Description, c.UpdatedAt, c.ID)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("category not found or no change")
	}
	return nil
}

func (r *categoryRepository) Delete(id int) error {
	query := `UPDATE categories SET deleted_at = $1 WHERE id = $2`
	_, err := r.db.Exec(query, time.Now(), id)
	return err
}
