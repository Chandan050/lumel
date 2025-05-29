package main

import (
	"log"

	"github.com/chandan050/lumel/internal/handler"
	"github.com/chandan050/lumel/internal/repository"
	"github.com/chandan050/lumel/internal/service"
	"github.com/chandan050/lumel/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loding .env file %v", err)
	}

	conn := ("host=localhost port=5432 user=postgres password=password dbname=sales sslmode=disable")
	dbName := repository.Connect(conn)

	logger := utils.InitLogger()

	srv := service.SalesService{DB: dbName, Logger: logger}
	handler.Svc = &srv
	RegisterRoutes(r)

	utils.RunMigrations(dbName)
	r.Run(":8080")

}

func RegisterRoutes(router *gin.Engine) {
	router.GET("/total-customers", handler.TotalCustomersHandler)
	router.GET("/total-orders", handler.TotalOrdersHandler)
	router.GET("/average-order-value", handler.AverageOrderValueHandler)
	router.POST("/refresh", handler.RefreshHandler)
}
