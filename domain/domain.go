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
	ID string `json:"id,omitempty"`
}

func (r ResourceID) String() string {
	return r.ID
}

func NewUser(name string) *User {
	return &User{ID: ResourceID{ID: uuid.NewString()}, Name: name}
}

func CreateUserCommand(u *User) error {
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
