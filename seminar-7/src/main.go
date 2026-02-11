package main

import (
	"fmt"
	"log"
	"seminar7/src/query"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := Connect()
	fmt.Println("Query 1:", query.Query1(db))
	fmt.Println("Query 2:", query.Query2(db))
	fmt.Println("Query 3:", query.Query3(db))
	fmt.Println("Query 4:", query.Query4(db))
	fmt.Println("Query 5:", query.Query5(db))
}

func Connect() *gorm.DB {
	dsn := "host=olympics user=olympics password=olympics dbname=olympics port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}
	return db
}
