package crud

import (
	"gorm-product/entity"

	"gorm.io/gorm"
)

func InsertOrder(db *gorm.DB, order *entity.Orders) {
	result := db.Create(order)

	if result.Error != nil {
		panic("failed to add data order")
	}
}

func InsertOrderDetails(db *gorm.DB, details *[]entity.OrderDetails) {
	result := db.Create(details)

	if result.Error != nil {
		panic("failed to add data orders")
	}
}

func ShowAllOrders(db *gorm.DB) []entity.OrderDetails {
	var orders []entity.OrderDetails

	result := db.Joins("join products as pd ON pd.kode_produk = order_details.kode_produk AND pd.stok > 90").Preload("Products").Joins("Orders").Find(&orders)

	if result.Error != nil {
		panic("failed show orders data")
	}
	return orders
}

func GetOneRow(db *gorm.DB, id_order int) entity.Orders {
	var order entity.Orders

	result := db.Where("id_order =?", id_order).First(&order)

	if result.Error != nil {
		panic("failed show data order")
	}

	return order
}

func UpdateOrder(db *gorm.DB, id_order int, data_order entity.Orders) {
	result := db.Where("id_order = ?", id_order).Updates(&data_order)

	if result.Error != nil {
		panic("failed update data order")
	}
}

func DeleteOrderDetail(db *gorm.DB, id_order_detail int) {
	result := db.Where("id_order_detail", id_order_detail).Delete(&entity.OrderDetails{})

	if result.Error != nil {
		panic("failed delete order detail")
	}
}
