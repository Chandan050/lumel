// internal/utils/csv_loader.go
package utils

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

func LoadCSVData(db *sqlx.DB, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	_, _ = reader.Read()

	tx := db.MustBegin()
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		// Insert into customers
		tx.MustExec(`INSERT INTO customers (id, name, email, address)
            VALUES ($1, $2, $3, $4)
            ON CONFLICT (id) DO NOTHING`,
			row[2], row[12], row[13], row[14])

		// Insert into products
		tx.MustExec(`INSERT INTO products (id, name, category,price)
            VALUES ($1, $2, $3)
            ON CONFLICT (id) DO NOTHING`,
			row[1], row[3], row[4], row[8])

		// Insert into orders
		date, _ := time.Parse("2006-01-02", row[6])
		quantity, _ := strconv.Atoi(row[7])
		unitPrice, _ := strconv.ParseFloat(row[8], 64)
		discount, _ := strconv.ParseFloat(row[9], 64)
		shippingCost, _ := strconv.ParseFloat(row[10], 64)

		tx.MustExec(`INSERT INTO orders (
            id, customer_id, region, sale_date,
            quantity, unit_price, discount, shipping_cost,
            payment_method, product_id)
            VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
            ON CONFLICT (id) DO NOTHING`,
			row[0], row[2], row[5], date, quantity,
			unitPrice, discount, shippingCost,
			row[11], row[1])
	}

	return tx.Commit()
}
