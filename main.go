package main

import (
	"deni1688/cqrsdemo/domain"
	"fmt"
)

func main() {
	john := domain.NewUser("John Doe")
	if err := domain.CreateUserCommand(john); err != nil {
		panic(err)
	}
	fmt.Printf("User created with id: %s\n", john.ID)

	john, err := domain.GetUserQuery(john.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User found with id: %s\n", john.ID)

	merry := domain.NewUser("Merry Doe")
	if err = domain.CreateUserCommand(merry); err != nil {
		panic(err)
	}
	fmt.Printf("User created with id: %s\n", merry.ID)

	merry, err = domain.GetUserQuery(merry.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User found with id: %s\n", merry.ID)
}
