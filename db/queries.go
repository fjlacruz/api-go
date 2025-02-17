package db

import (
	"database/sql"
	"gin-api/models"
	"log"
)

// CreateProduct inserta un nuevo producto en la base de datos.
func CreateProduct(p *models.Product) error {
	_, err := DB.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", p.Name, p.Price)
	return err
}

// GetProduct obtiene un producto por su ID.
func GetProduct(id int) (*models.Product, error) {
	var p models.Product
	err := DB.QueryRow("SELECT id, name, price FROM products WHERE id = $1", id).
		Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

// GetAllProducts obtiene todos los productos.
func GetAllProducts() ([]models.Product, error) {
	rows, err := DB.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			log.Println("Error al escanear producto:", err)
			continue
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

// UpdateProduct actualiza un producto existente.
func UpdateProduct(p *models.Product) error {
	_, err := DB.Exec("UPDATE products SET name = $1, price = $2 WHERE id = $3", p.Name, p.Price, p.ID)
	return err
}

// DeleteProduct elimina un producto por su ID.
func DeleteProduct(id int) error {
	_, err := DB.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}
