package crud

import (
	"gorm-product/entity"

	"gorm.io/gorm"
)

func InsertProduct(db *gorm.DB, product *entity.Products) {
	result := db.Create(product)

	if result.Error != nil {
		panic("failed to add data product")
	}
}

func ShowAllProducts(db *gorm.DB) []entity.Products {
	var produk []entity.Products

	result := db.Where("kode_produk = ?", "001").Find(&produk)

	if result.Error != nil {
		panic("failed show product data")
	}
	return produk
}

func GetOneRow(db *gorm.DB, kode_produk string) entity.Products {
	var product entity.Products

	db.First(&product)

	return product
}

func UpdateProduct(db *gorm.DB, kode_produk string, data_produk entity.Products) {
	result := db.Where("kode_produk = ?", kode_produk).Updates(&data_produk)

	if result.Error != nil {
		panic("failed update product data")
	}
}

func DeleteProduct(db *gorm.DB, kode_produk string) {
	result := db.Delete(&entity.Products{
		KodeProduk: kode_produk,
	})

	if result.Error != nil {
		panic("failed update data product")
	}
}
