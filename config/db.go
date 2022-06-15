package config

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
	"c2/models"
	"fmt"
)

func DBInit() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=go-c2 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to cennect to database")
	}
	fmt.Println("Success connect to DB using GORM")
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Item{})
	return db
}