package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Rafipratama22/go_market/entity"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func SetupDatabase() *gorm.DB {
	godotenv.Load()
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_host := os.Getenv("DB_HOST")

	fmt.Println("DB_NAME: ", db_name, "DB_USER: ", db_user, "DB_PASS: ", db_pass, "DB_HOST: ", db_host)
	db, err := gorm.Open("mysql", "sprout:gr0wingDlab@tcp(localhost:3306)/go_market?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})
	return db
}

func CloseDatabase(db *gorm.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func Setup() {
	db := SetupDatabase()
	defer db.Close()
}

func SetupMongoDatabase() (*mongo.Client) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client
}

func GetCollection(db *mongo.Client, schema string) (*mongo.Collection) {
	collection := db.Database("local").Collection(schema)
	return collection
} 