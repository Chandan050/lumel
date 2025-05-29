package models

import "time"

type Order struct {
	ID            string    `json:"id"`
	CustomerID    string    `json:"customer_id"`
	Region        string    ` json:"region"`
	SaleDate      time.Time `json:"sale_date"`
	Quantity      int       ` json:"quantity"`
	UnitPrice     float64   `json:"unit_price"`
	Discount      float64   `json:"discount"`
	ShippingCost  float64   ` json:"shipping_cost"`
	PaymentMethod string    ` json:"payment_method"`
	ProductID     string    `json:"product_id"`
}
