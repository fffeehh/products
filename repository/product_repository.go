package repository

import (
	"attempt/models"
	"database/sql"
	"fmt"
	"log"
)


type ProductRepository interface{
	Init() error
	AddProduct(product models.Product) (models.Product, error)
	GetProducts() ([]models.Product, error)
	GetProductByID(id int) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(id int) error
}

type sqliteProductRepository struct {
	db *sql.DB
}

func NewSQLitePrroductRepository(db *sql.DB) ProductRepository {
	return &sqliteProductRepository{db: db}
}

func (r *sqliteProductRepository) Init() error {
	createTableSQL := `
	CREATE TABLE IF NOT EXIST products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
		numberOfGram TEXT
	);`

	_, err := r.db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create products table: %w", err)
	}
	fmt.Println("table was created successfully")
	return nil
}


