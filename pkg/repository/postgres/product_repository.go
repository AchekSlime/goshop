package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"goshop"
)

type ProductsPostgresRepository struct {
	db *sqlx.DB
}

func NewProductsPostgresRepository(db *sqlx.DB) *ProductsPostgresRepository {
	return &ProductsPostgresRepository{db: db}
}

func (p ProductsPostgresRepository) Create(product *goshop.Product) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	var productId int
	createProductQuery := fmt.Sprintf("INSERT INTO %s (title, price, holder_name) values ($1, $2, $3) RETURNING id", ProductsTable)

	row := tx.QueryRow(createProductQuery, product.Title, product.Price, product.HolderName)
	err = row.Scan(&productId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return productId, tx.Commit()
}

func (p ProductsPostgresRepository) Delete(productId int) {
	//TODO implement me
	panic("implement me")
}

func (p ProductsPostgresRepository) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (p ProductsPostgresRepository) GetById(productId int) {
	//TODO implement me
	panic("implement me")
}

func (p ProductsPostgresRepository) GetAfter(productId int, count int) {
	//TODO implement me
	panic("implement me")
}

func (p ProductsPostgresRepository) GetFirst(count int) ([]goshop.Product, error) {
	var products []goshop.Product
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.price, ti.holder_name FROM %s ti  LIMIT $1`,
		ProductsTable)

	if err := p.db.Get(&products, query, count); err != nil {
		return products, err
	}

	return products, nil
}
