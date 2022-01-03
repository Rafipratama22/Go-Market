package main

import (
	// "github.com/Rafipratama22/go_market/config"
	"github.com/Rafipratama22/go_market/routes"
	gorm "github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
	routing routes.Server
)

func main() {
	route := routing.Router()
	defer db.Close()
	route.Run(":8080")
}