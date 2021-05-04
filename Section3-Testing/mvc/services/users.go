package services

import (
	"github.com/nwneisen/MicroservicesInGo/Section2-MVC/mvc/domain"
)

type usersService struct{}

var (
	UsersService usersService
)

func (u *usersService) GetUser(userId int64) (*domain.User, error) {
	return domain.UsersDomain.GetUser(userId)
}
