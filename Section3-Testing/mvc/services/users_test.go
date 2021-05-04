package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUsersNotFoundInDatabase(t *testing.T) {
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}
