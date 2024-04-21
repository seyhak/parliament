package logic

import "testing"

func TestUpdateAnswerState(t *testing.T) {
	keyEx := "k1"
	t.Run("Test answer index is 2", func(t *testing.T) {
		InitiateState(nil)
		state := GetGlobalState()
		state.updateAnswerState(keyEx, 2)
		values := state.ProblemState.Votes[keyEx]
		for idx, val := range values {
			if idx == 1 {
				if val != 1 {
					t.Errorf("Value should be 1, got %v", val)
				}
			} else {
				if val != 0 {
					t.Errorf("Value should be 0, got %v", val)
				}
			}
		}
	})
	t.Run("Test updating multiple values twice", func(t *testing.T) {
		InitiateState(nil)
		state := GetGlobalState()
		state.updateAnswerState(keyEx, 2)
		state.updateAnswerState(keyEx, 2)

		values := state.ProblemState.Votes[keyEx]
		for idx, val := range values {
			if idx == 1 {
				if val != 2 {
					t.Errorf("Value should be 2, got %v", val)
				}
			} else {
				if val != 0 {
					t.Errorf("Value should be 0, got %v", val)
				}
			}
		}
	})
}

func TestAddProblemToHistory(t *testing.T) {
	t.Run("Test problem is copied", func(t *testing.T) {
		InitiateState(nil)
		state := GetGlobalState()
		answers := []string{"Ans 1", "Ans2"}
		problem := Problem{"aa", "title", "some content", answers}
		state.addProblemToHistory(problem)
		historyProblem := state.ProblemState.HistoryProblems["aa"]
		if historyProblem.Content != problem.Content {
			t.Error("History not copied properly")
		}
		if historyProblem.Answers[0] != problem.Answers[0] {
			t.Error("History not copied properly")
		}

		historyProblem.Answers[0] = "new value 123"
		historyProblem.Content = "new Content 315"

		if historyProblem.Content == problem.Content {
			t.Error("History not copied properly")
		}
		if historyProblem.Answers[0] == problem.Answers[0] {
			t.Error("History not copied properly")
		}
	})
	t.Run("Add problem to history in state", func(t *testing.T) {
		InitiateState(nil)
		state := GetGlobalState()
		answers := []string{"Ans 1", "Ans2"}
		problem := Problem{"aa", "title", "some content", answers}
		state.addProblemToHistory(problem)
		if len(state.ProblemState.HistoryProblems) != 1 {
			t.Error("Failed to add problem to history")
		}
		if state.ProblemState.HistoryProblems["aa"].Id != problem.Id {
			t.Error("Adding not proper id")
		}
		if state.ProblemState.HistoryProblems["aa"].Title != problem.Title {
			t.Error("Adding not proper Title")
		}
		if state.ProblemState.HistoryProblems["aa"].Content != problem.Content {
			t.Error("Adding not proper Content")
		}
		for idx, val := range answers {
			if state.ProblemState.HistoryProblems["aa"].Answers[idx] != val {
				t.Error("Adding not proper Answers")
			}
		}
	})
}

func TestMain(m *testing.M) {
	InitiateState(nil)
	m.Run()
}
