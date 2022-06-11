package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"goshop/pkg/entities"
	"goshop/pkg/repository"
	"log"
	"strconv"
)

const (
	ProductTable  = "product"
	CategoryTable = "category"
)

func NewPostgresDb(cfg repository.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DbName,
		cfg.Password,
	))
	if err != nil {
		return nil, fmt.Errorf("can't connect to postgres db")
	}
	return db, nil
}

func InitCategories(categoryRepository *CategoryRepository, productRepository *ProductRepository) error {
	names := []string{"Фитнес", "Футбол", "Баскетбол", "Волейбол", "Одежда"}
	for i := 0; i < len(names); i++ {
		category := entities.Category{
			Name: names[i],
		}
		id, err := categoryRepository.Create(&category)
		if err != nil {
			log.Println("db: can't create category")
			return err
		}
		category.Id = id
		InitProducts(productRepository, &category)
	}

	return nil
}

func InitProducts(productRepository *ProductRepository, category *entities.Category) error {
	for i := 0; i < 15; i++ {
		product := entities.Product{
			Title:      category.Name + " " + strconv.Itoa(i+1),
			Price:      i*100 - i*25,
			HolderName: "Stabbers",
			CategoryId: category.Id,
		}
		id, err := productRepository.Create(&product)
		if err != nil {
			log.Println("db: can't create product: " + err.Error())
			return err
		}
		product.Id = id
	}

	return nil
}
