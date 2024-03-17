package library

import (
	"database/sql"
	"fmt"

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

type order struct {
	ID           int    `json:"id"`
	CustomerName string `json:"customerName"`
	ProductId    int    `json:"productId"`
	ProductName  string `json:"productName"`
	ProductPrice string `json:"productPrice"`
	Quantity     int    `json:"quantity"`
	Total        int    `json:"total"`
	OrderStatus  string `json:"orderStatus"`
}

func getOrders(db *sql.DB) ([]order, error) {
	rows, err := db.Query("SELECT o.id, o.customerName, p.name, p.price, oi.quantity, o.Total, o.Status FROM orders o, products p, order_items oi WHERE o.id = oi.order_id AND oi.product_id = p.id")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []order{}

	for rows.Next() {
		var o order

		if err := rows.Scan(&o.ID, &o.CustomerName, &o.ProductName, &o.ProductPrice, &o.Quantity, &o.Total, &o.OrderStatus); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func (o *order) createOrder(db *sql.DB) error {
	res, err := db.Exec("INSERT INTO orders(customerName, total, status) VALUES(?, ?, ?)", o.CustomerName, o.Total, o.OrderStatus)

	if err != nil {
		fmt.Println("createOrder first error: ", err.Error())
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println("createOrder second error: ", err.Error())
		return err
	}

	orderId := int(id)

	res1, err := db.Exec("INSERT INTO order_items(order_id, product_id, quantity) VALUES(?, ?, ?)", orderId, o.ProductId, o.Quantity)

	if err != nil {
		fmt.Println("createOrder third error: ", err.Error())
		return err
	}
	id1, err := res1.LastInsertId()

	if err != nil {
		fmt.Println("createOrder fourth error: ", err.Error())
		return err
	}

	fmt.Println("Inserted order_items id: ", id1)
	return nil
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

func (p *product) createProduct(db *sql.DB) error {
	res, err := db.Exec("INSERT INTO products(productCode, name, inventory, price, status) VALUES (?,?,?,?,?)", p.ProductCode, p.Name, p.Inventory, p.Price, p.Status)

	if err != nil {
		fmt.Println("createProduct err1: ", err.Error())
		return err
	}
	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println("createProduct err2: ", err.Error())
		return err
	}
	p.ID = int(id)
	return nil
}
