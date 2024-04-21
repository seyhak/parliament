package logic

import (
	"fmt"

	logic_types "github.com/seyhak/parliament/logic/types"
)

func presentResult() {
	fmt.Println("---result---")
	fmt.Println(GetGlobalState().ProblemState.Votes)
	state := GetGlobalState()
	for key, votes := range state.ProblemState.Votes {
		currentProblem := state.ProblemState.HistoryProblems[key]
		fmt.Printf("Answers for %s \n", currentProblem.Title)
		for idx, vote := range votes {
			if len(currentProblem.Answers) <= idx {
				break
			}
			fmt.Printf("Votes for: \"%s\" - %v\n", currentProblem.Answers[idx], vote)
		}
		fmt.Println("")
	}
}
func checkResult() int {
	// TODO
	return 0

}
func HandleResult() {
	presentResult()
	checkResult()
}

func RunParliament(users *[]logic_types.User) {
	handleProblems()
	HandleResult()
}
