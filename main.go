package main

import (
	"database/sql"
	"fmt"
	"log"

	"example.com/library"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	id        int
	name      string
	inventory int
	price     int
}

func main() {
	db, err := sql.Open("sqlite3", "./practiceit.db")

	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := db.Query("SELECT id, name, inventory, price FROM products")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var p Product

		rows.Scan(&p.id, &p.name, &p.inventory, &p.price)

		fmt.Println(p.id, " ", p.name, " ", p.inventory, " ", p.price)
	}

	app := library.App{}
	app.Port = ":8545"

	// Initialize DB connection
	app.Initialize()

	app.Run()
}
