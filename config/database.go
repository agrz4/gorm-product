package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB() {
	host := "localhost"
	username := ""
	password := ""
	database := "goprod"
	port := "5432"

	dsn := " host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + "sslmode=disable TimeZone=Asia/Jakarta"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})

	if err != nil {
		panic("Can't connect to database")
	}
}
