package logic

import (
	"encoding/json"
	"log"
)

// TODO create config file
const ConfigMaxChoicesSize = 6

type UserState struct {
	SkipUser bool
}

type Votes = map[string][]int
type ProblemState struct {
	Votes map[string][]int // key is uuid of a problem, second is array of Votes per answer
}

type State struct {
	UserState    UserState
	ProblemState ProblemState
}

var defaultUserState = UserState{
	SkipUser: true,
}

var defaultProblemState = ProblemState{
	Votes: make(Votes),
}

var defaultState = State{
	defaultUserState,
	defaultProblemState,
}

var globalState *State = nil

func InitiateState(state *State) *State {
	if state != nil {
		globalState = state
		return globalState
	}
	data, _ := json.Marshal(defaultState)

	var newState State
	err := json.Unmarshal(data, &newState)
	if err != nil {
		log.Println("error copying state")
	}

	globalState = &newState
	return globalState
}

func GetGlobalState() *State {
	return globalState
}

func (state *State) updateAnswerState(key string, answerIndex int) {
	modifiedIdx := answerIndex - 1
	if state.ProblemState.Votes[key] == nil {
		state.ProblemState.Votes[key] = make([]int, ConfigMaxChoicesSize)
	}
	state.ProblemState.Votes[key][modifiedIdx]++
}
