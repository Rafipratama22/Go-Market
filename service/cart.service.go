package service

import (
	"github.com/Rafipratama22/go_market/repository"
	"github.com/Rafipratama22/go_market/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartService interface {
	CreateCart(cart schema.Cart) *mongo.InsertOneResult
	FindAll() []bson.M
	FindOne(order_id string) bson.M
	UpdateOne(order_id string, body schema.Cart) *mongo.UpdateResult
	DeleteOne(order_id string) *mongo.DeleteResult
}

type cartService struct {
	cartRepo repository.CartRepository
}

func NewCartService(repository repository.CartRepository) CartService {
	return &cartService{
		cartRepo: repository,
	}
}

func (c *cartService) CreateCart(cart schema.Cart) *mongo.InsertOneResult{
	return c.cartRepo.CreateCart(cart)
}

func (c *cartService) FindAll() []bson.M{
	return c.cartRepo.FindAll()
}

func (c *cartService) FindOne(order_id string) bson.M {
	return c.cartRepo.FindOne(order_id)
}

func (c *cartService) UpdateOne(order_id string, body schema.Cart) *mongo.UpdateResult {
	return c.cartRepo.UpdateOne(order_id, body)
}

func (c *cartService) DeleteOne(order_id string) *mongo.DeleteResult {
	return c.cartRepo.DeleteOne(order_id)
}