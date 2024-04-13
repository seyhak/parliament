package logic

import (
	"fmt"

	logic_types "github.com/seyhak/parliament/logic/types"
)

func CreateUser() logic_types.User {
	fmt.Println("Please provide username:")
	state := GetGlobalState()
	if state.UserState.SkipUser {
		return logic_types.User{Name: "Default user", Party: 0}
	}
	var username string
	fmt.Scanln(&username)
	user := logic_types.User{Name: username, Party: 0}
	return user
}
