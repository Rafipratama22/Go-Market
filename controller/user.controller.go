package controller

import (
	"fmt"
	"strconv"

	"github.com/Rafipratama22/go_market/entity"
	"github.com/Rafipratama22/go_market/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	UserFindId(ctx *gin.Context)
	UserFindAll(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) CreateUser(ctx *gin.Context) {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(user)
	c.userService.CreateUser(user)
	ctx.JSON(200, gin.H{
		"message": "User created successfully",
	})
}

func (c *userController) UserFindId(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println("error:", err)
	}
	userData := c.userService.FindUserId(userId)
	ctx.JSON(200, gin.H{
		"data": userData,
	})
}

func (c *userController) UserFindAll(ctx *gin.Context) {
	usersData := c.userService.FindAll()
	ctx.JSON(200, gin.H{
		"data": usersData,
	})
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	var user entity.User
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println("error:", err)
	}
	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(user)
	c.userService.UpdateUser(user, userId)
	ctx.JSON(200, gin.H{
		"message": "User updated successfully",
	})
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println("error:", err)
	}
	c.userService.DeleteUser(userId)
	ctx.JSON(200, gin.H{
		"message": "User deleted successfully",
	})
}
