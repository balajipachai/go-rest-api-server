package library

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type product struct {
	ID          int    `json:"id"`
	ProductCode string `json:"productCode"`
	Name        string `json:"name"`
	Inventory   int    `json:"inventory"`
	Price       int    `json:"price"`
	Status      string `json:"status"`
}

// The `getProducts` function retrieves product data from a database and returns it as a slice of
// product structs.
func getProducts(db *sql.DB) ([]product, error) {
	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []product{}
	for rows.Next() {
		var p product

		if err := rows.Scan(&p.ID, &p.ProductCode, &p.Name, &p.Inventory, &p.Price, &p.Status); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// The `func (p *product) getProduct(db *sql.DB) error` method is a method associated with the
// `product` struct in Go. It is a method that allows you to retrieve a specific product from the
// database based on the product's ID.
func (p *product) getProduct(db *sql.DB) error {
	return db.QueryRow("SELECT productCode, name, inventory, price, status FROM products WHERE id = ?", p.ID).Scan(&p.ProductCode, &p.Name, &p.Inventory, &p.Price, &p.Status)
}
