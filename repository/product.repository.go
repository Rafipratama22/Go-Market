package repository

import (
	"github.com/Rafipratama22/go_market/entity"
	"github.com/Rafipratama22/go_market/helper"
	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ProductRepository interface {
	SaveProduct(product entity.Product)
	GetProduct(productId int) entity.Product
	GetProducts(pagination map[string]string, condition map[string]interface{}) []entity.Product
	UpdateProduct(product entity.Product, productId int)
	DeleteProduct(productId int)
	BulkCreateProduct()
}

type productRepository struct {
	db *gorm.DB
}

type productStruct struct {
	Products []entity.Product `json:"products"`
}

func NewProductRepo(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func filterProduct(query map[string]interface{}) func (db *gorm.DB)*gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		modelProduct := db.Model(&entity.Product{})
		
		fmt.Println(query)
		if query["name"] != nil {
			modelProduct = modelProduct.Where("name LIKE ?", "%" + query["name"].(string) + "%")
		}
		if query["price"] != nil {
			modelProduct = modelProduct.Where("price = ?", query["price"])
		}
		if query["category_id"] != nil  {
			modelProduct = modelProduct.Where("category_id IN ?", query["category_id"])
		}
		if query["stock"] != "" {
			fmt.Println()
			modelProduct = modelProduct.Where("stock > ?", 0)
		}
		if query["department_id"] != nil {
			modelProduct = modelProduct.Where("department_id IN ?", query["department_id"])
		}
		if query["max_price"] != 0 {
			modelProduct = modelProduct.Where("price < ?", query["max_price"])
		}
		if query["min_price"] != 0 {
			modelProduct = modelProduct.Where("price > ?", query["min_price"])
		}
		return modelProduct
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

func (c *productRepository) GetProducts(pagination map[string]string, condition map[string]interface{}) []entity.Product {
	var products []entity.Product
	c.db.Scopes(filterProduct(condition), helper.Paginate(pagination)).Find(&products)
	return products
}

func (c *productRepository) UpdateProduct(product entity.Product, productId int) {
	c.db.Model(&product).Where("id = ?", productId).Updates(&product)
}

func (c *productRepository) DeleteProduct(productId int) {
	var product entity.Product
	c.db.First(&product, productId)
	c.db.Delete(&product)
}

func (c *productRepository) BulkCreateProduct() {
	var product entity.Product
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened products.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var products productStruct

	json.Unmarshal(byteValue, &products)
	for i := 29878; i < len(products.Products); i++ {
		fmt.Println(i)
		c.db.Model(&product).Create(&products.Products[i])
	}
}