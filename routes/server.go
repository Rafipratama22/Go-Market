package routes

import (
	"github.com/Rafipratama22/go_market/config"
	"github.com/Rafipratama22/go_market/controller"
	"github.com/Rafipratama22/go_market/repository"
	"github.com/Rafipratama22/go_market/service"
	"github.com/gin-gonic/gin"
	gorm "github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	router *gin.Engine
}

var (
	db                *gorm.DB                     = config.SetupDatabase()
	mongodb           *mongo.Client                = config.SetupMongoDatabase()
	productRepo       repository.ProductRepository = repository.NewProductRepo(db)
	productService    service.ProductService       = service.NewProductService(productRepo)
	productController controller.ProductController = controller.NewProductController(productService)
	userRepo          repository.UserRepository    = repository.NewUserRepo(db)
	userService       service.UserService          = service.NewUserService(userRepo)
	userController    controller.UserController    = controller.NewUserController(userService)
	cartRepo	repository.CartRepository = repository.NewCartRepo(mongodb)
	cartService service.CartService = service.NewCartService(cartRepo)
	cartController controller.CartController = controller.NewCartController(cartService)

)

func MainServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server) Router() *gin.Engine {
	route := gin.Default()
	var apiName string = "/api/v1"

	productRoute := route.Group(apiName + "/product")
	{
		productRoute.GET("/", func(c *gin.Context) {
			productController.GetProducts(c)
		})
		productRoute.POST("/", func(c *gin.Context) {
			productController.CreateProduct(c)
		})
		productRoute.GET("/:id", func(c *gin.Context) {
			productController.GetProductID(c)
		})
		productRoute.PUT("/:id", func(c *gin.Context) {
			productController.UpdateProduct(c)
		})
		productRoute.DELETE("/:id", func(c *gin.Context) {
			productController.DeleteProduct(c)
		})
	}

	userRoute := route.Group(apiName + "/user")
	{
		userRoute.GET("/", func(c *gin.Context) {
			userController.UserFindAll(c)
		})
		userRoute.POST("/", func(c *gin.Context) {
			userController.CreateUser(c)
		})
		userRoute.GET("/:id", func(c *gin.Context) {
			userController.UserFindId(c)
		})
		userRoute.PUT("/:id", func(c *gin.Context) {
			userController.UpdateUser(c)
		})
		userRoute.DELETE("/:id", func(c *gin.Context) {
			userController.DeleteUser(c)
		})
	}

	cartRoute := route.Group(apiName + "/cart")
	{
		cartRoute.GET("/", func (ctx *gin.Context) {
			cartController.FindAllCart(ctx)
		})
		cartRoute.POST("/", func (ctx *gin.Context) {
			cartController.CreateCart(ctx)
		})
		cartRoute.GET("/:order_id", func (ctx *gin.Context) {
			cartController.FindOneCart(ctx)
		})
		cartRoute.PUT("/:order_id", func (ctx *gin.Context) {
			cartController.UpdateOneCart(ctx)
		})
		cartRoute.DELETE("/:order_id", func (ctx *gin.Context) {
			cartController.DeleteOneCart(ctx)
		})
	}

	return route
}
