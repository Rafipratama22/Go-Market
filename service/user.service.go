package service

import (
	"github.com/Rafipratama22/go_market/entity"
	"github.com/Rafipratama22/go_market/repository"
)

type UserService interface {
	CreateUser(user entity.User)
	FindUserId(userId int) entity.User
	FindAll() []entity.User
	UpdateUser(user entity.User, userId int)
	DeleteUser(userId int)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		userRepo: repository,
	}
}

func (c *userService) CreateUser(user entity.User) {
	c.userRepo.CreateUser(user)
}

func (c *userService) FindUserId(userId int) entity.User {
	return c.userRepo.FindUserId(userId)
}

func (c *userService) FindAll() []entity.User {
	return c.userRepo.FindAll()
}

func (c *userService) UpdateUser(user entity.User, userId int) {
	c.userRepo.UpdateUser(user, userId)
}

func (c *userService) DeleteUser(userId int) {
	c.userRepo.DeleteUser(userId)
}
