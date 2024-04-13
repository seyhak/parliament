package logic

import (
	"fmt"

	logic_types "github.com/seyhak/parliament/logic/types"
)

func checkResult() {
	fmt.Println(GetGlobalState().ProblemState.votes)
}

func RunParliament(users *[]logic_types.User) {
	handleProblems()
	// TODO
	checkResult()
}
