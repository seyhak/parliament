package main

import (
	"fmt"

	logic "github.com/seyhak/parliament/logic"
	logic_types "github.com/seyhak/parliament/logic/types"
	tanks "github.com/seyhak/parliament/tanks"
)

const (
	TANKS = iota
	PARLIAMENT
)

func runParliament() {
	fmt.Println("Welcome to parliament!")
	logic.InitiateState(nil)
	user := logic.CreateUser()
	fmt.Println(user.Name)

	users := []logic_types.User{user}
	logic.RunParliament(&users)
}

func main() {
	const gameToRun = TANKS
	switch gameToRun {
	case TANKS:
		tanks.RunTanks()
	default:
		runParliament()
	}
}
