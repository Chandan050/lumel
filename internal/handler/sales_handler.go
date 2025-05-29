package handler

import (
	"net/http"
	"time"

	"github.com/chandan050/lumel/internal/service"
	"github.com/gin-gonic/gin"
)

var Svc *service.SalesService

// Utility function to parse dates
func parseDateRange(c *gin.Context) (time.Time, time.Time, bool) {
	start := c.Query("start")
	end := c.Query("end")

	from, err1 := time.Parse("2006-01-02", start)
	to, err2 := time.Parse("2006-01-02", end)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return time.Time{}, time.Time{}, false
	}

	return from, to, true
}

// GET /total-customers
func TotalCustomersHandler(c *gin.Context) {
	from, to, ok := parseDateRange(c)
	if !ok {
		return
	}

	count, err := Svc.GetTotalCustomers(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_customers": count})
}

// GET /total-orders
func TotalOrdersHandler(c *gin.Context) {
	from, to, ok := parseDateRange(c)
	if !ok {
		return
	}

	count, err := Svc.GetTotalOrders(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_orders": count})
}

// GET /average-order-value
func AverageOrderValueHandler(c *gin.Context) {
	from, to, ok := parseDateRange(c)
	if !ok {
		return
	}

	avg, err := Svc.GetAverageOrderValue(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"average_order_value": avg})
}

// POST /refresh
func RefreshHandler(c *gin.Context) {
	err := Svc.RefreshData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "refresh successful"})
}
