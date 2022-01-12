package repository

import (
	"github.com/Rafipratama22/go_market/entity"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/Rafipratama22/go_market/middleware"
)

type UserRepository interface {
	CreateUser(user entity.User)
	FindUserId(userId int) entity.User
	FindAll() []entity.User
	UpdateUser(user entity.User, userId int)
	DeleteUser(userId int)
	LoginUser(user entity.User) (string, bool)
	FindName(name string) entity.User
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (c *userRepository) CreateUser(user entity.User) {
	c.db.Create(&user)
}

func (c *userRepository) FindUserId(userId int) entity.User {
	var user entity.User
	c.db.First(&user, userId)
	return user
}

func (c *userRepository) FindAll() []entity.User {
	var users []entity.User
	c.db.Find(&users)
	return users
}

func (c *userRepository) UpdateUser(user entity.User, userId int) {
	c.db.Model(&user).Where("id = ?", userId).Update(&user)
}

func (c *userRepository) DeleteUser(userId int) {
	var user entity.User
	c.db.First(&user, userId)
	c.db.Delete(&user)
}

func (c *userRepository) LoginUser(user entity.User) (string, bool) {
	var userLogin entity.User
	fmt.Println(user.Email)
	c.db.Where("email = ?", user.Email).First(&userLogin)
	fmt.Println(userLogin)
	if userLogin.Email != "" {
		usered, err := middleware.CreateToken(userLogin.Name)
		fmt.Println(usered, err)
		if err == nil {
			return usered, false
		} else {
			return "", true
		}
	} else {
		return "", true
	}
}

func (c *userRepository) FindName(name string) entity.User {
	var user entity.User
	c.db.Where("name = ?", name).First(&user)
	return user
}