package crud

import (
	"fmt"
	"gorm-product/config"
	"gorm-product/entity"
	"testing"
)

func TestShowOrder(t *testing.T) {
	config.ConnectDB()

	orders := ShowAllOrders(config.DB)

	for _, val := range orders {
		fmt.Println(val)
	}
}

func TestUpdateOrder(t *testing.T) {
	config.ConnectDB()
	id_order := 1
	data_order := entity.Orders{
		PaymentMethod: "Bank Transfer",
		Total:         50000,
	}

	UpdateOrder(config.DB, id_order, data_order)
	order := GetOneRow(config.DB, id_order)

	fmt.Println("New data order")
	fmt.Println(order)

	if order.PaymentMethod == "Cash" && order.Total == 30000 {
		t.Errorf("failed update data order")
	}
}

func TestDeleteOrderDetails(t *testing.T) {
	config.ConnectDB()
	DeleteOrderDetail(config.DB, 2)
}
