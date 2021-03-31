package services

import (
	"github.com/nwneisen/MicroservicesInGo/Section2-MVC/mvc/domain"
)

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
