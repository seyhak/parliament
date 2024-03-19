package main

import (
	"fmt"

	"github.com/seyhak/parliament/types"

	logic "github.com/seyhak/parliament/logic"
)

func main() {
	fmt.Println("Welcome to parliament!")

	fmt.Println("!")
	user := logic.CreateUser()
	fmt.Println(user.Name)

	users := []types.User{user}
	logic.RunParliament(&users)
}
