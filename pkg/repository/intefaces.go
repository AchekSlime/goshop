package repository

import (
	"goshop/pkg/entities"
)

type Config struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
}

type ProductRepositoryInterface interface {
	Create(p *entities.Product) (int, error)
	Delete(productId int)
	GetAll()
	GetById(productId int) (*entities.Product, error)
	GetAfterByCategory(categoryId, productId int, limit int) ([]entities.Product, error)
	GetFirst(count int) ([]entities.Product, error)
}

type CategoryRepositoryInterface interface {
	Create(c *entities.Category) (int, error)
	GetAll() ([]entities.Category, error)
	GetById(categoryId int) (*entities.Category, error)
}

type CartRepositoryInterface interface {
	Create(productId, userId int) error
	GetProductsId(userId int) ([]int, error)
	Delete(productId, userId int) error
}

type StoryRepositoryInterface interface {
	Create(pId []int, userId int) (int, error)
	CreateMaping(storyId, productId int) error
	GetAllStoryId(userId int) ([]int, []string, error)
	GetProductsId(storyId int) ([]int, error)
}

type Storage struct {
	ProductRepositoryInterface
	CategoryRepositoryInterface
	CartRepositoryInterface
	StoryRepositoryInterface
}
