package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UsersDomain.GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err, "we were expecting an error when user id is 0")
}

func TestGetUserUserFound(t *testing.T) {
	user, err := UsersDomain.GetUser(123)

	assert.NotNil(t, user, "we were expecting a user with id 123")
	assert.Nil(t, err, "we were not expecting an error when user id is 123")
}
