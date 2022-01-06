package controller

import (
	"fmt"
	"strconv"

	"github.com/Rafipratama22/go_market/entity"
	"github.com/Rafipratama22/go_market/service"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	CreateProduct(ctx *gin.Context)
	GetProductID(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	BulkCreateProduct(ctx *gin.Context)
}

type productController struct {
	productService service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &productController{
		productService: service,
	}
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	var product entity.Product
	err := ctx.ShouldBindJSON(&product)
	fmt.Println(ctx.PostForm("name"))
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(product)
	c.productService.CreateProduct(product)
	ctx.JSON(200, gin.H{
		"message": "Product created successfully",
	})
}

func (c *productController) GetProductID(ctx *gin.Context) {
	productId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println("error:", err)
	}
	productData := c.productService.GetProductID(productId)
	ctx.JSON(200, gin.H{
		"data": productData,
	})
}

func (c *productController) GetProducts(ctx *gin.Context) {
	pagination := make(map[string]interface{})
	pagination["page"] = ctx.Query("page")
	pagination["limit"] = ctx.Query("limit")
	productsData := c.productService.GetProducts(pagination)
	ctx.JSON(200, gin.H{
		"data": productsData,
	})
}

func (c *productController) UpdateProduct(ctx *gin.Context) {
	var product entity.Product
	productId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println("error:", err)
	}
	err = ctx.ShouldBindJSON(&product)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(product)
	c.productService.UpdateProduct(product, productId)
	ctx.JSON(200, gin.H{
		"message": "Product updated successfully",
	})
}

func (c *productController) DeleteProduct(ctx *gin.Context) {
	productId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println("error:", err)
	}
	c.productService.DeleteProduct(productId)
	ctx.JSON(200, gin.H{
		"message": "Product deleted successfully",
	})
}

func (c *productController) BulkCreateProduct(ctx *gin.Context) {
	c.productService.BulkCreateProduct()
	ctx.JSON(200, gin.H{
		"message": "Product bulk created successfully",
	})
}

// func (c *productController) ListCatalog(ctx *gin.Context) {

// }