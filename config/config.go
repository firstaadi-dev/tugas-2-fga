package config

import (
	"assignment-dua-fga/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	DSN := "root:password@tcp(127.0.0.1:3306)/fga?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Order{})
	return db
}
