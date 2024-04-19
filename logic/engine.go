package logic

import (
	"fmt"

	logic_types "github.com/seyhak/parliament/logic/types"
)

func checkResult() {
	fmt.Println("---result---")
	fmt.Println(GetGlobalState().ProblemState.Votes)
}

func RunParliament(users *[]logic_types.User) {
	handleProblems()
	// TODO
	checkResult()
}
