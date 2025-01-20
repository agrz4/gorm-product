package entity

import "time"

type Products struct {
	KodeProduk string `gorm:"column:kode_produk;primaryKey"`
	NamaProduk string `gorm:"column:nama_produk"`
	Stok       int    `gorm:"column:stok"`
	Harga      int    `gorm:"column:harga"`
}

type Orders struct {
	IdOrder       int       `gorm:"column:id_order;primaryKey;autoIncrement"`
	TanggalOrder  time.Time `gorm:"column:tanggal_order"`
	PaymentMethod string    `gorm:"column:payment_method"`
	Total         int       `gorm:"column:total"`
}

type OrderDetails struct {
	IdOrderDetail      int      `gorm:"column:id_order_detail;primaryKey;autoIncrement"`
	OrdersIdOrder      int      `gorm:"column:id_order"`
	Orders             Orders   `gorm:"foreignKey:OrdersIdOrder"`
	ProductsKodeProduk string   `gorm:"column:kode_produk"`
	Products           Products `gorm:"foreignKey:ProductsKodeProduk"`
	Qty                int      `gorm:"column:qty"`
}
