package database

import (
	"database/sql"

	"github.com/xtuser777/goapi/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{
		db: db,
	}
}

func (pd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(
			&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL,
		); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (pd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product
	err := pd.db.QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(
		&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pd *ProductDB) GetProductsByCategory(categoryId string) ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT * FROM products WHERE category_id = ?", categoryId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(
			&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL,
		); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (pd *ProductDB) CreateProduct(product *entity.Product) (*entity.Product, error) {
	_, err := pd.db.Exec(
		"INSERT INTO products(id, name, description, price, image_url, category_id) VALUES(?, ?, ?, ?, ?, ?)",
		product.ID, product.Name, product.Description, product.Price, product.ImageURL, product.CategoryID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
