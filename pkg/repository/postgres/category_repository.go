package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"goshop/pkg/entities"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (p CategoryRepository) Create(category *entities.Category) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	var categoryId int
	createProductQuery := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", CategoryTable)

	row := tx.QueryRow(createProductQuery, category.Name)
	err = row.Scan(&categoryId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return categoryId, tx.Commit()
}

func (p CategoryRepository) GetAll() ([]entities.Category, error) {
	var categories []entities.Category
	query := fmt.Sprintf(`SELECT c.id, c.name FROM %s c`, CategoryTable)

	if err := p.db.Select(&categories, query); err != nil {
		return nil, err
	}

	return categories, nil
}

func (p CategoryRepository) GetById(id int) (*entities.Category, error) {
	var category []entities.Category
	query := fmt.Sprintf(`SELECT c.id, c.name FROM %s c WHERE c.id = $1`, CategoryTable)

	if err := p.db.Select(&category, query, id); err != nil {
		return nil, err
	}

	return &category[0], nil
}
