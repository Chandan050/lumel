package utils

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func RunMigrations(db *sqlx.DB) {
	schema := `
CREATE TABLE IF NOT EXISTS customers (
	id TEXT PRIMARY KEY,
	name TEXT,
	email TEXT,
	address TEXT
);

CREATE TABLE IF NOT EXISTS products (
	id TEXT PRIMARY KEY,
	name TEXT,
	category TEXT
);

CREATE TABLE IF NOT EXISTS orders (
	id TEXT PRIMARY KEY,
	product_id TEXT,
	customer_id TEXT,
	region TEXT,
	sale_date DATE,
	quantity INTEGER,
	unit_price FLOAT,
	discount FLOAT,
	shipping_cost FLOAT,
	payment_method TEXT,
	FOREIGN KEY (product_id) REFERENCES products(id),
	FOREIGN KEY (customer_id) REFERENCES customers(id)
);
`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
