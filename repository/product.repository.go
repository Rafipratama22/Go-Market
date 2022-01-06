package repository

import (
	"github.com/Rafipratama22/go_market/entity"
	"github.com/Rafipratama22/go_market/helper"
	"github.com/jinzhu/gorm"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ProductRepository interface {
	SaveProduct(product entity.Product)
	GetProduct(productId int) entity.Product
	GetProducts(query map[string]interface{}) []entity.Product
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
		var products []entity.Product
		return db.Where(query).Find(&products)
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

func (c *productRepository) GetProducts(query map[string]interface{}) []entity.Product {
	var products []entity.Product
	// c.db.Find(&products)
	// return products
	// return filterProduct(query, c.db)
	queryPage := make(map[string]string)
	queryPage["page"] = query["page"].(string)
	queryPage["limit"] = query["limit"].(string)
	c.db.Scopes(filterProduct(query), helper.Paginate(queryPage)).Find(&products)
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