package tanks

import (
	"fmt"
	"math/rand"
)

const MAX_WIDTH = 1000
const MAX_HEIGHT = 700
const MAX_LVL_DIFF = 10

func getNextStep(prevStepHeight int) int {
	var nextStepSign int
	willBeHigherThanBottom := prevStepHeight-MAX_LVL_DIFF > 0
	if willBeHigherThanBottom {
		nextStepSign = 1 * (rand.Intn(2) * -1)
	} else {
		nextStepSign = 1
	}
	diff := rand.Intn(MAX_LVL_DIFF)
	return prevStepHeight + (diff * nextStepSign)
}

func generateMap() {
	board := make([][]bool, MAX_WIDTH)
	line := make([]int, MAX_WIDTH)

	firstStep := rand.Intn(2 * MAX_HEIGHT / 3)
	for idx, _ := range line {
		nextStep := getNextStep(firstStep)
		fmt.Println(idx, nextStep)
		line[idx] = nextStep
	}

	for x := range MAX_WIDTH {
		board[x] = make([]bool, MAX_HEIGHT)
		maxHeight := line[x]
		for y := range MAX_HEIGHT {
			isBelowOrEqMaxH := y <= maxHeight
			board[x][y] = isBelowOrEqMaxH
		}
		// println(board[x]) TODO
	}
	// fmt.Print(board)
}
