package repository

import (
	// "github.com/Rafipratama22/go_market/config"
	"context"
	"fmt"

	"github.com/Rafipratama22/go_market/schema"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CartRepository interface {
	CreateCart(cart schema.Cart) *mongo.InsertOneResult
	FindAll() []bson.M
	FindOne(order_id string) bson.M
	UpdateOne(order_id string, body schema.Cart) *mongo.UpdateResult
	DeleteOne(order_id string) *mongo.DeleteResult
}

type cartRepository struct {
	db *mongo.Client
}


func NewCartRepo(db *mongo.Client) CartRepository {
	return &cartRepository{
		db: db,
	}
}

func (c *cartRepository) CreateCart(cart schema.Cart) *mongo.InsertOneResult {
	collection := c.db.Database("Go-Market").Collection("Carts")
	res, err := collection.InsertOne(context.TODO(), &cart)
	if err != nil {
		panic(err)
	}
	return res
}

func (c *cartRepository) FindAll() []bson.M {
	var results []bson.M
	findOptions := options.Find()
	collection := c.db.Database("Go-Market").Collection("Carts")
	cursor, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	if results == nil {
		panic("error in cursor all")
	}
	fmt.Println(results)
	return results
}

func (c *cartRepository) FindOne(order_id string) bson.M {
	var results bson.M
	// findOptions := options.Find()
	var objId primitive.ObjectID
	if order_id != "" {
		objId, _ = primitive.ObjectIDFromHex(order_id)
	}
	collection := c.db.Database("local").Collection("Carts")
	if err := collection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&results); err != nil {
		panic(err)
	}
	return results
}

func (c *cartRepository) UpdateOne(order_id string, body schema.Cart) *mongo.UpdateResult {
	var objId primitive.ObjectID
	if order_id != "" {
		objId, _ = primitive.ObjectIDFromHex(order_id)
	}
	collection := c.db.Database("local").Collection("Carts")
	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{
		"$set": &body,
	}); 
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return result
}

func (c *cartRepository) DeleteOne(order_id string) *mongo.DeleteResult{
	var objId primitive.ObjectID
	if order_id != "" {
		objId, _ = primitive.ObjectIDFromHex(order_id)
	}
	collection := c.db.Database("local").Collection("Carts")
	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objId})
	if err != nil {
		panic(err)
	}
	return res
} 