package domain

import (
	"errors"
	"github.com/google/uuid"
)

var userStore = map[ResourceID]User{}

type User struct {
	ID   ResourceID `json:"id,omitempty"`
	Name string     `json:"name,omitempty"`
}

type ResourceID struct {
	Value string
}

func (r ResourceID) String() string {
	return r.Value
}

func NewUser(name string) *User {
	return &User{ID: ResourceID{Value: uuid.NewString()}, Name: name}
}

func RegisterUserCommand(u *User) error {
	if u.ID.String() == "" {
		return errors.New("user id is empty")
	}

	if _, ok := userStore[u.ID]; ok {
		return errors.New("user already exists")
	}

	userStore[u.ID] = *u

	return nil
}

func GetUserQuery(id ResourceID) (*User, error) {
	if id.String() == "" {
		return nil, errors.New("user id is empty")
	}

	if u, ok := userStore[id]; ok {
		return &u, nil
	}

	return nil, errors.New("user not found")
}
