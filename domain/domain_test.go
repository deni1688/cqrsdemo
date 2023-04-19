package domain

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceID_String(t *testing.T) {
	id := ResourceID{Value: "test-id"}

	assert.Equal(t, "test-id", id.String())
}

func TestNewUser(t *testing.T) {
	user := NewUser("John Doe")

	assert.NotNil(t, user.ID.String())
	assert.Equal(t, "John Doe", user.Name)
}

func TestRegisterUserCommand(t *testing.T) {
	t.Run("register valid user", func(t *testing.T) {
		user := NewUser("Jane Doe")
		err := RegisterUserCommand(user)

		assert.Nil(t, err)
	})

	t.Run("register user with empty ID", func(t *testing.T) {
		user := &User{Name: "Empty ID"}
		err := RegisterUserCommand(user)

		assert.NotNil(t, err)
		assert.Equal(t, errors.New("user id is empty"), err)
	})

	t.Run("register already existing user", func(t *testing.T) {
		user := NewUser("John Smith")
		_ = RegisterUserCommand(user)
		err := RegisterUserCommand(user)

		assert.NotNil(t, err)
		assert.Equal(t, errors.New("user already exists"), err)
	})
}

func TestGetUserQuery(t *testing.T) {
	t.Run("get existing user", func(t *testing.T) {
		user := NewUser("Alice")
		_ = RegisterUserCommand(user)
		retrievedUser, err := GetUserQuery(user.ID)

		assert.Nil(t, err)
		assert.Equal(t, user, retrievedUser)
	})

	t.Run("get user with empty ID", func(t *testing.T) {
		id := ResourceID{Value: ""}
		retrievedUser, err := GetUserQuery(id)

		assert.Nil(t, retrievedUser)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("user id is empty"), err)
	})

	t.Run("get non-existing user", func(t *testing.T) {
		id := ResourceID{Value: "non-existing-id"}
		retrievedUser, err := GetUserQuery(id)

		assert.Nil(t, retrievedUser)
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("user not found"), err)
	})
}
