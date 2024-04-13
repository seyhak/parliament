package main

import (
	"fmt"

	logic "github.com/seyhak/parliament/logic"
	logic_types "github.com/seyhak/parliament/logic/types"
)

func main() {
	fmt.Println("Welcome to parliament!")
	logic.InitiateState(nil)
	user := logic.CreateUser()
	fmt.Println(user.Name)

	users := []logic_types.User{user}
	logic.RunParliament(&users)
}
