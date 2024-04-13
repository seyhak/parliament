package logic

// TODO create config file
const ConfigMaxChoicesSize = 6

type UserState struct {
	SkipUser bool
}

type votes = map[string][]int
type ProblemState struct {
	votes map[string][]int // key is uuid of a problem, second is array of votes per answer
}

type State struct {
	UserState    UserState
	ProblemState ProblemState
}

var defaultUserState = UserState{
	SkipUser: true,
}

var defaultProblemState = ProblemState{
	votes: make(votes),
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
	globalState = &defaultState
	return globalState
}

func GetGlobalState() *State {
	return globalState
}

func (state *State) updateAnswerState(key string, answerIndex int) {
	if state.ProblemState.votes[key] == nil {
		state.ProblemState.votes[key] = make([]int, ConfigMaxChoicesSize)
	}
	// prevValue := state.ProblemState.votes[key][answerIndex]
	state.ProblemState.votes[key][answerIndex]++
}
