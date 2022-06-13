package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type StoryRepository struct {
	db *sqlx.DB
}

func NewStoryRepository(db *sqlx.DB) *StoryRepository {
	return &StoryRepository{db: db}
}

func (c *StoryRepository) Create(pId []int, userId int) (int, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return 0, err
	}

	createProductQuery := fmt.Sprintf("INSERT INTO %s (user_id, creation_date) values ($1, $2) RETURNING id", StoryTable)

	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal(err)
	}

	var sId int
	row := tx.QueryRow(createProductQuery, userId, time.Now().In(location))
	err = row.Scan(&sId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}

	for _, v := range pId {
		if err := c.CreateMaping(sId, v); err != nil {
			return 0, err
		}
	}
	return sId, nil
}

func (c *StoryRepository) CreateMaping(storyId, productId int) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	createProductQuery := fmt.Sprintf("INSERT INTO %s (story_id, product_id) values ($1, $2)", StoryProductTable)

	tx.QueryRow(createProductQuery, storyId, productId)
	return tx.Commit()
}

func (c *StoryRepository) GetAllStoryId(userId int) ([]int, []string, error) {
	var sId []int
	query := fmt.Sprintf(`SELECT s.id FROM %s s WHERE s.user_id = $1`,
		StoryTable)

	if err := c.db.Select(&sId, query, userId); err != nil {
		log.Println("db: GetAllStoryId error: " + err.Error())
		return nil, nil, err
	}

	var dates []string
	query = fmt.Sprintf(`SELECT s.creation_date FROM %s s WHERE s.user_id = $1`,
		StoryTable)

	if err := c.db.Select(&dates, query, userId); err != nil {
		log.Println("db: GetAllStoryId error: " + err.Error())
		return nil, nil, err
	}

	return sId, dates, nil
}

func (c *StoryRepository) GetProductsId(storyId int) ([]int, error) {
	var pId []int
	query := fmt.Sprintf(`SELECT s.product_id FROM %s s WHERE s.story_id = $1`,
		StoryProductTable)

	if err := c.db.Select(&pId, query, storyId); err != nil {
		log.Println("db: GetProductsId error: " + err.Error())
		return nil, err
	}

	return pId, nil
}
