package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type car struct {
	ID       int64
	Brand    string
	MaxSpeed int64
	FuelType string
}

// docker run -d -p 54322:5432 -e POSTGRES_PASSWORD=postgres postgres
func main() {
	dsn := "host=localhost port=54322 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err := sqlDb.Ping(); err != nil {
		panic(err)
	}
	db.AutoMigrate(&car{})
	defer db.Model(&car{}).Delete(&car{}, "1 = 1")
	db.Create([]*car{
		{Brand: "Ferrari", MaxSpeed: 250, FuelType: "gasoline"},
		{Brand: "FIAT", MaxSpeed: 180, FuelType: "gasoline"},
	})
	var cars []car
	if err := db.Model(&car{}).Find(&cars).Error; err != nil {
		panic(err)
	}
	for _, v := range cars {
		fmt.Println(v)
	}
}
