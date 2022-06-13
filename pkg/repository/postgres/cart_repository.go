package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type CartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (c *CartRepository) Create(productId, userId int) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	createProductQuery := fmt.Sprintf("INSERT INTO %s (product_id, user_id) values ($1, $2)", CartTable)

	tx.QueryRow(createProductQuery, productId, userId)
	return tx.Commit()
}

func (c *CartRepository) GetProductsId(userId int) ([]int, error) {
	var productsId []int
	query := fmt.Sprintf(`SELECT c.product_id FROM %s c WHERE c.user_id = $1`,
		CartTable)

	if err := c.db.Select(&productsId, query, userId); err != nil {
		log.Println("db: GetProducts error: " + err.Error())
		return nil, err
	}

	return productsId, nil
}

func (c *CartRepository) Delete(productId, userId int) error {
	query := fmt.Sprintf(`DELETE FROM %s c WHERE product_id = $1 AND user_id = $2`,
		CartTable)

	if _, err := c.db.Exec(query, productId, userId); err != nil {
		log.Println("db: GetProducts error: " + err.Error())
		return err
	}

	return nil
}
