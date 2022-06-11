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
	GetById(productId int)
	GetAfterByCategory(categoryId, productId int, limit int) ([]entities.Product, error)
	GetFirst(count int) ([]entities.Product, error)
}

type CategoryRepositoryInterface interface {
	Create(c *entities.Category) (int, error)
	GetAll() ([]entities.Category, error)
	GetById(categoryId int) (*entities.Category, error)
}

type Storage struct {
	ProductRepositoryInterface
	CategoryRepositoryInterface
}
