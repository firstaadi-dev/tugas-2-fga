package main

import (
	"assignment-dua-fga/config"
	"assignment-dua-fga/order/delivery/http"
	"assignment-dua-fga/order/repository/mysql"
	"assignment-dua-fga/order/usecase"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	conn := config.InitDB()
	r := gin.Default()

	orderRepo := mysql.NewMysqlOrderRepository(conn)
	orderUcase := usecase.NewOrderUsecase(orderRepo)

	http.NewOrderHandler(r, orderUcase)
	err := r.Run()
	if err != nil {
		os.Exit(1)
	}
}
