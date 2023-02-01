package database

import (
	"fmt"
	"log"

	"github.com/v3rxn/go-short/config"
	"github.com/v3rxn/go-short/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySQLUser, config.MySQLPassword, config.DBHost, config.MySQLDB)
	
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database!")
	}

	log.Println("Connected to database")

	DB = connection

	DB.AutoMigrate(&models.Link{})

}