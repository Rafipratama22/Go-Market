package service

import (
	"github.com/Rafipratama22/go_market/entity"
	"github.com/Rafipratama22/go_market/repository"
)

type ProductService interface {
	CreateProduct(product entity.Product)
	GetProductID(productId int) entity.Product
	GetProducts() []entity.Product
	UpdateProduct(product entity.Product, productId int)
	DeleteProduct(productId int)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) ProductService {
	return &productService{
		productRepo: repository,
	}
}

func (c *productService) CreateProduct(product entity.Product) {
	c.productRepo.SaveProduct(product)
}

func (c *productService) GetProductID(productId int) entity.Product {
	return c.productRepo.GetProduct(productId)
}

func (c *productService) GetProducts() []entity.Product {
	return c.productRepo.GetProducts()
}

func (c *productService) UpdateProduct(product entity.Product, productId int) {
	c.productRepo.UpdateProduct(product, productId)
}

func (c *productService) DeleteProduct(productId int) {
	c.productRepo.DeleteProduct(productId)
}
