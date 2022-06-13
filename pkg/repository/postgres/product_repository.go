package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"goshop/pkg/entities"
	"log"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p ProductRepository) Create(product *entities.Product) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	var productId int
	createProductQuery := fmt.Sprintf("INSERT INTO %s (title, price, holder_name, category_id) values ($1, $2, $3, $4) RETURNING id", ProductTable)

	row := tx.QueryRow(createProductQuery, product.Title, product.Price, product.HolderName, product.CategoryId)
	err = row.Scan(&productId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return productId, tx.Commit()
}

func (p ProductRepository) Delete(productId int) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) GetById(productId int) (*entities.Product, error) {
	var products []entities.Product
	query := fmt.Sprintf(`SELECT p.id, p.title, p.price, p.holder_name, p.category_id FROM %s p WHERE p.id = $1`,
		ProductTable)

	if err := p.db.Select(&products, query, productId); err != nil {
		log.Println("db: GetAfterByCategory error: " + err.Error())
		return nil, err
	}

	return &products[0], nil
}

func (p ProductRepository) GetAfterByCategory(categoryId, productId int, limit int) ([]entities.Product, error) {
	var products []entities.Product
	query := fmt.Sprintf(`SELECT p.id, p.title, p.price, p.holder_name, p.category_id FROM %s p WHERE p.category_id = $1 AND p.id > $2 LIMIT $3`,
		ProductTable)

	if err := p.db.Select(&products, query, categoryId, productId, limit); err != nil {
		log.Println("db: GetAfterByCategory error: " + err.Error())
		return nil, err
	}

	return products, nil
}

func (p ProductRepository) GetFirst(count int) ([]entities.Product, error) {
	var products []entities.Product
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.price, ti.holder_name FROM %s ti LIMIT $1`,
		ProductTable)

	if err := p.db.Select(&products, query, count); err != nil {
		return products, err
	}

	return products, nil
}
