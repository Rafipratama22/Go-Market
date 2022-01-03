package repository

import (
	"github.com/Rafipratama22/go_market/entity"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User)
	FindUserId(userId int) entity.User
	FindAll() []entity.User
	UpdateUser(user entity.User, userId int)
	DeleteUser(userId int)
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
