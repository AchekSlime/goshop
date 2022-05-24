package repository

import (
	"goshop"
)

type Config struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
}

//type Repository interface {
//	New(cfg Config)
//	Create(p *Product) error
//	Delete(productId int)
//	GetAll()
//	GetById(productId int)
//	GetAfter(productId int, count int)
//	GetFirst(count int)
//}

type ProductRepository interface {
	Create(p *goshop.Product) (int, error)
	Delete(productId int)
	GetAll()
	GetById(productId int)
	GetAfter(productId int, count int)
	GetFirst(count int) ([]goshop.Product, error)
}

type Repository struct {
	ProductRepository
}
