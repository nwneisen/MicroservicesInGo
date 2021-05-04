package domain

import (
	"fmt"
	"log"
)

var (
	users = map[int64]*User{
		123: {
			Id:        123,
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "FooBar@email.com",
		},
	}

	UsersDomain usersDomain
)

type usersDomain struct{}

type User struct {
	Id        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (u *usersDomain) GetUser(userId int64) (*User, error) {
	log.Println("accessing the database")
	user := users[userId]
	if user == nil {
		return nil, fmt.Errorf("User Id %v was not found", userId)
	}

	return user, nil
}
