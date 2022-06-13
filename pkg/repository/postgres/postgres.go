package postgres

import (
	"bufio"
	"fmt"
	"github.com/jmoiron/sqlx"
	"goshop/pkg/entities"
	"goshop/pkg/repository"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ProductTable      = "product"
	CategoryTable     = "category"
	CartTable         = "cart"
	StoryTable        = "story"
	StoryProductTable = "story_product"
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
	names := ReadDirs()

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

func ReadDirs() []string {
	s := make([]string, 0)
	items, _ := ioutil.ReadDir("./images/")
	for _, item := range items {
		if item.IsDir() {
			s = append(s, item.Name())
		}
	}
	return s
}

func InitProducts(productRepository *ProductRepository, category *entities.Category) error {
	files := ReadFiles(category.Name)

	for _, v := range files {
		itemName := strings.Split(v, ";")[0]
		price, _ := strconv.Atoi(strings.Split(v, ";")[1])

		product := entities.Product{
			Title:      itemName,
			Price:      price,
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

func ReadFiles(dirName string) []string {
	images := make([]string, 0)
	file, err := os.Open("./images/" + dirName + "/config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		images = append(images, scanner.Text())
	}

	return images
}
