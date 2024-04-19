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

func TestMain(m *testing.M) {
	InitiateState(nil)
	m.Run()
}
