package domain

import (
	"fmt"
)

var (
	items = map[int64]*Item{
		123: {
			Id:          123,
			Description: "Stuff and things",
		},
	}
)

type itemsDomain struct{}

var (
	ItemsDomain itemsDomain
)

type Item struct {
	Id          uint64 `json:"id"`
	Description string `json:"description"`
}

func (i *itemsDomain) GetItem(itemId int64) (*Item, error) {
	item := items[itemId]
	if item == nil {
		return nil, fmt.Errorf("Item Id %v was not found", itemId)
	}

	return item, nil
}
