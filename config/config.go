package config

import (
	"context"
	"os"
	"time"

	"github.com/Rafipratama22/go_market/entity"
	// gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	// "gorm.io/gorm/schema"
)

func SetupDatabase() *gorm.DB {
	godotenv.Load()
	// db_name := os.Getenv("DB_NAME")
	// db_user := os.Getenv("DB_USER")
	// db_pass := os.Getenv("DB_PASS")
	// db_host := os.Getenv("DB_HOST")
	// db_port := os.Getenv("DB_PORT")

	// fmt.Println("DB_NAME: ", db_name, "DB_USER: ", db_user, "DB_PASS: ", db_pass, "DB_HOST: ", db_host)
	// db, err := gorm.Open("mysql", db_user + ":" + db_pass + "@tcp(" + db_host + ":"+ db_port +")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local")
	// db, err := gorm.Open("mysql", "sprout:gr0wingDlab@tcp(localhost:3306)/go_market?charset=utf8mb4&parseTime=True&loc=Local")
	// db, err := gorm.Open("postgres", "host=db.jaiyputsidbjlpyullim.supabase.co user=postgres dbname=postgres password=rafipratama port=6543")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=postgres password=a dbname=postgres sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	// postgres://postgres:[YOUR-PASSWORD]@db.jaiyputsidbjlpyullim.supabase.co:6543/postgres
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Address{}, &entity.PaymentTransaction{}, &entity.Order{}, &entity.OrderDetail{}, &entity.Kecamatan{}, &entity.Kelurahan{}, &entity.City{}, &entity.Province{})
	//
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans.Environment = midtrans.Sandbox
	return db
}

func Setup() {
	// db := SetupDatabase()
	// defer db.Close()
}

func SetupMongoDatabase() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://rafipratama:rafipratama22@devgomarket.fmpo9.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
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

func GetCollection(db *mongo.Client, schema string) *mongo.Collection {
	collection := db.Database("local").Collection(schema)
	return collection
}
