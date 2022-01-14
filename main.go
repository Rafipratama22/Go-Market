package main

import (
	// "github.com/Rafipratama22/go_market/config"
	"log"

	"github.com/Rafipratama22/go_market/routes"
	// gorm "github.com/jinzhu/gorm"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
	routing routes.Server
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}	

func main() {
	route := routing.Router()
	// defer db.Close()
	route.Run(":8080")
}