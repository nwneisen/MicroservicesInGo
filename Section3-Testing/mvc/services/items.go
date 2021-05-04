package services

import (
	"errors"

	"github.com/nwneisen/MicroservicesInGo/Section3-Testing/mvc/domain"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (i *itemsService) GetItem(itemId string) (*domain.Item, error) {
	return nil, errors.New("implement me")
}
