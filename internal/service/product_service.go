package service

import (
	"github.com/xtuser777/goapi/internal/entity"

	"github.com/xtuser777/goapi/internal/database"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(ProductDB database.ProductDB) *ProductService {
	return &ProductService{
		ProductDB: ProductDB,
	}
}

func (cs *ProductService) GetProducts() ([]*entity.Product, error) {
	categories, err := cs.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cs *ProductService) GetProductsByCategory(categoryID string) ([]*entity.Product, error) {
	categories, err := cs.ProductDB.GetProductsByCategory(categoryID)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cs *ProductService) CreateProduct(name, description, imageURL, categoryID string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, imageURL, categoryID, price)
	_, err := cs.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (cs *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := cs.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
