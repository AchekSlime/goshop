package main

import (
	_ "github.com/lib/pq"
	"goshop/pkg/repository"
	"goshop/pkg/repository/postgres"
	"goshop/pkg/telegram"
	"log"
)

func main() {
	// создать конфиг postgres
	cfg := repository.Config{
		Host:     "localhost",
		Port:     "5435",
		DbName:   "goshop",
		User:     "postgres",
		Password: "qwe",
	}

	// создать бд
	db, err := postgres.NewPostgresDb(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	// создать базы данных
	productRepository := postgres.NewProductRepository(db)
	categoryRepository := postgres.NewCategoryRepository(db)
	cartRepository := postgres.NewCartRepository(db)
	storyRepository := postgres.NewStoryRepository(db)

	//postgres.InitCategories(categoryRepository, productRepository)

	// создать репозиторий
	repository := repository.Storage{
		ProductRepositoryInterface:  productRepository,
		CategoryRepositoryInterface: categoryRepository,
		CartRepositoryInterface:     cartRepository,
		StoryRepositoryInterface:    storyRepository,
	}

	// создать бота
	bot := telegram.NewBot("5301308275:AAExkTKVknQJH8YfRfQHiukG-B8MX3vtybY", repository)

	bot.Start()
}
