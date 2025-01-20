package main

import (
	"gorm-product/config"
	crud_order "gorm-product/crud/orders"
	crud_product "gorm-product/crud/products"
	"gorm-product/entity"
	"time"
)

func main() {
	config.ConnectDB()

	product := entity.Products{
		KodeProduk: "001",
		NamaProduk: "Indomie",
		Stok:       100,
		Harga:      3000,
	}

	crud_product.InsertProduct(config.DB, &product)
	product = entity.Products{
		KodeProduk: "002",
		NamaProduk: "Sushi",
		Stok:       150,
		Harga:      5000,
	}

	crud_product.InsertProduct(config.DB, &product)

	order := entity.Orders{
		TanggalOrder:  time.Now(),
		PaymentMethod: "Cash",
		Total:         8000,
	}

	crud_order.InsertOrder(config.DB, &order)

	details := []entity.OrderDetails{
		{
			Orders: order,
			Products: entity.Products{
				KodeProduk: "001",
			},
			Qty: 10,
		},
		{
			Orders: order,
			Products: entity.Products{
				KodeProduk: "002",
			},
			Qty: 5,
		},
	}
	crud_order.InsertOrderDetails(config.DB, &details)
}
