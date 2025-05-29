package service

import (
	"log"
	"time"

	"github.com/chandan050/lumel/utils"
	"github.com/jmoiron/sqlx"
)

type SalesService struct {
	DB     *sqlx.DB
	Logger *log.Logger
}

// service com registery
func (s *SalesService) GetTotalCustomers(from, to time.Time) (int, error) {
	var count int
	err := s.DB.Get(&count, `
		SELECT COUNT(DISTINCT customer_id) 
		FROM orders 
		WHERE sale_date BETWEEN $1 AND $2`, from, to)
	return count, err
}

func (s *SalesService) GetTotalOrders(from, to time.Time) (int, error) {
	var count int
	err := s.DB.Get(&count, `
		SELECT COUNT(*) 
		FROM orders 
		WHERE sale_date BETWEEN $1 AND $2`, from, to)
	return count, err
}

func (s *SalesService) GetAverageOrderValue(from, to time.Time) (float64, error) {
	var avg float64
	err := s.DB.Get(&avg, `
		SELECT COALESCE(AVG(unit_price * quantity), 0)
		FROM orders
		WHERE sale_date BETWEEN $1 AND $2`, from, to)
	return avg, err
}

func (s *SalesService) RefreshData() error {
	err := utils.LoadCSVData(s.DB, "csv/sales_data.csv")
	if err != nil {
		s.Logger.Println("Failed to refresh:", err)
		return err
	}
	s.Logger.Println("Refresh successful")
	return nil

}
