package crud

import (
	"fmt"
	"gorm-product/config"
	"gorm-product/entity"
	"testing"
)

func TestShowProduct(t *testing.T) {
	config.ConnectDB()

	produk := ShowAllProducts(config.DB)

	for _, val := range produk {
		fmt.Println(val)
	}
}

func TestShowOneRow(t *testing.T) {
	config.ConnectDB()
	kode_produk := "001"
	produk := GetOneRow(config.DB, kode_produk)

	if produk.KodeProduk == "" {
		t.Errorf("Data :" + kode_produk + " empty")
	}
	fmt.Println(produk)
}

func TestUpdateProduk(t *testing.T) {
	config.ConnectDB()

	kode_produk := "001"
	data_produk_update := entity.Products{
		NamaProduk: "Indomie Rebus",
		Stok:       1000,
		Harga:      2500,
	}

	UpdateProduct(config.DB, kode_produk, data_produk_update)

	produk := GetOneRow(config.DB, kode_produk)

	fmt.Println("New product data")
	fmt.Println(produk)

	if produk.NamaProduk == "Indomie" && produk.Stok == 100 && produk.Harga == 3000 {
		t.Errorf("failed update product data" + kode_produk)
	}
}

func TestDeleteProduct(t *testing.T) {
	config.ConnectDB()
	DeleteProduct(config.DB, "002")
}
