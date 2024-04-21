package logic

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
)

func openWholeFile(directory *string) string {
	// read whole file
	dat, err := os.ReadFile(*directory)
	if err != nil {
		log.Fatalln(err)
	}
	return string(dat)
}

func constructProblemFromFile(path *string) Problem {
	fileContent := openWholeFile(path)
	var problem Problem
	error := json.Unmarshal([]byte(fileContent), &problem)
	if error != nil {
		fmt.Println(error)
	}

	return problem
}

func getAbsoluteProblemPath(problemName *string) *string {
	const relDir = "./assets/problems"
	var pathWithProblem string
	if problemName == nil {
		pathWithProblem = relDir
	} else {
		pathWithProblem = fmt.Sprintf("%s/%s", relDir, *problemName)
	}
	absDir, err := filepath.Abs(pathWithProblem)
	if err != nil {
		log.Println("Error getting absolute path:", err)
		return nil
	}
	return &absDir
}

func loadProblems() []fs.FileInfo {
	absDir := getAbsoluteProblemPath(nil)
	// Open the directory
	directory, err := os.Open(*absDir)
	if err != nil {
		log.Println("Error opening directory:", err)
		return nil
	}
	defer directory.Close()

	// Read the directory entries
	entries, err := directory.Readdir(-1) // -1 means no limit
	if err != nil {
		log.Println("Error reading directory:", err)
		return nil
	}

	// Print all file names
	for _, entry := range entries {
		log.Println(entry.Name())
	}
	return entries
}

func (problem *Problem) presentProblem() {
	fmt.Printf("\n%s\n", problem.Title)
	fmt.Printf("Message: %s\n", problem.Content)
	fmt.Println("What do you want to do?")
	for idx, answer := range problem.Answers {
		fmt.Printf("%d: %s\n", idx+1, answer)
	}

}

func (problem *Problem) isAnswerValid(answer int) bool {
	availableOptions := make([]int, len(problem.Answers))
	for idx := range availableOptions {
		availableOptions[idx] = idx + 1
	}
	isValid := slices.Contains(availableOptions, answer)
	if !isValid {
		fmt.Println("Provide valid answer")
	}
	return isValid
}

func (problem *Problem) processAnswer() {
	var answerIndex int
	continueValidating := true
	for continueValidating {
		fmt.Scanln(&answerIndex)
		continueValidating = !problem.isAnswerValid(answerIndex)
	}
	fmt.Printf("You chose: %v\n\n", answerIndex)
	state := GetGlobalState()
	state.updateAnswerState(problem.Id, answerIndex)
}

func processProblem(file fs.FileInfo) {
	fileName := file.Name()
	absDir := getAbsoluteProblemPath(&fileName)
	// Open the directory
	fileContent, err := os.Open(*absDir)
	if err != nil {
		log.Println("Error opening file:", err)
		panic(err)
	}
	log.Println("Opening file:", fileName)
	defer fileContent.Close()

	problem := constructProblemFromFile(absDir)
	problem.presentProblem()
	problem.processAnswer()
	GetGlobalState().addProblemToHistory(problem)
}

func handleProblems() {
	problems := loadProblems()
	if problems == nil {
		return
	}
	for _, problem := range problems {
		log.Println("loading problem", problem.Name())
		processProblem(problem)
	}
}
