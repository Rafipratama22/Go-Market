package repository

import (
	"github.com/Rafipratama22/go_market/entity"
	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	SaveProduct(product entity.Product)
	GetProduct(productId int) entity.Product
	GetProducts() []entity.Product
	UpdateProduct(product entity.Product, productId int)
	DeleteProduct(productId int)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (c *productRepository) SaveProduct(product entity.Product) {
	c.db.Create(&product)
}

func (c *productRepository) GetProduct(productId int) entity.Product {
	var product entity.Product
	c.db.First(&product, productId)
	return product
}

func (c *productRepository) GetProducts() []entity.Product {
	var products []entity.Product
	c.db.Find(&products)
	return products
}

func (c *productRepository) UpdateProduct(product entity.Product, productId int) {
	c.db.Model(&product).Where("id = ?", productId).Update(&product)
}

func (c *productRepository) DeleteProduct(productId int) {
	var product entity.Product
	c.db.First(&product, productId)
	c.db.Delete(&product)
}
