package logic

import (
	"fmt"

	"github.com/seyhak/parliament/types"
)

func CreateUser() types.User {
	var username string
	fmt.Println("Please provide username:")
	fmt.Scanln(&username)
	user := types.User{Name: username, Party: 0}
	return user
}
