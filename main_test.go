package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCaseReport(t *testing.T) {
	InputStdin := `
REPORT
PLACE 1,1,NORTH
REPORT
`
	processLines(strings.NewReader(InputStdin))

	_, err := processCommands([]string{"MOVE"})
	fmt.Println(err)
	if err != nil {
		t.Error("Error position is invalid!")
	}
}
func TestPositiveCasePlace(t *testing.T) {
	InputStdin := `
LEFT
`
	processLines(strings.NewReader(InputStdin))

	_, err := processCommands([]string{"PLACE", "1,1,NORTH"})
	fmt.Println(err)
	if err != nil {
		t.Error("Error position is invalid!")
	}

}
func TestNegativeCasePlace1(t *testing.T) {
	InputStdin := `
LEFT
`
	processLines(strings.NewReader(InputStdin))

	_, err := processCommands([]string{"PLACE", "5,4,NORTH"})
	fmt.Println(err)
	if err == nil {
		t.Error("Error position is invalid!")
	}

}

func TestNegativeCasePlace2(t *testing.T) {
	InputStdin := `
LEFT
`
	processLines(strings.NewReader(InputStdin))

	_, err := processCommands([]string{"PLACE", "A,4,NORTH"})
	fmt.Println(err)
	if err == nil {
		t.Error("Error position is invalid!")
	}

}

func TestNegativeCasePlace3(t *testing.T) {
	InputStdin := `
LEFT
`
	processLines(strings.NewReader(InputStdin))

	_, err := processCommands([]string{"PLACE", "3,B,NORTH"})
	fmt.Println(err)
	if err == nil {
		t.Error("Error position is invalid!")
	}

}

func TestNegativeCasePlace4(t *testing.T) {
	InputStdin := `
LEFT
`
	processLines(strings.NewReader(InputStdin))

	_, err := processCommands([]string{"PLACE", "3,3,NORT"})
	fmt.Println(err)
	if err == nil {
		t.Error("Error position is invalid!")
	}

}

func TestNegativeCaseLeft(t *testing.T) {
	InputStdin := `
PLACE 1,1,NORTH
LEFT
MOVE
MOVE
MOVE
`
	processLines(strings.NewReader(InputStdin))

	_, err := processCommands([]string{"MOVE"})
	fmt.Println(err)
	if err == nil {
		t.Error("Error position is invalid!")
	}

}
func TestPositiveCaseLeft(t *testing.T) {
	InputStdin := `
PLACE 4,4,NORTH
LEFT
MOVE
MOVE
`
	processLines(strings.NewReader(InputStdin))

	_, err := processCommands([]string{"MOVE"})
	fmt.Println(err)
	if err != nil {
		t.Error("Error position is invalid!")
	}

}

func TestNegativeCaseRight(t *testing.T) {
	InputStdin := `
PLACE 1,1,NORTH
RIGHT
MOVE
MOVE
MOVE
`
	processLines(strings.NewReader(InputStdin))

	_, err := processCommands([]string{"MOVE"})
	fmt.Println(err)
	if err == nil {
		t.Error("Error position is invalid!")
	}

}
func TestPositiveCaseRight(t *testing.T) {
	InputStdin := `
PLACE 1,1,NORTH
RIGHT
MOVE
MOVE
`
	processLines(strings.NewReader(InputStdin))

	_, err := processCommands([]string{"MOVE"})
	fmt.Println(err)
	if err != nil {
		t.Error("Error position is invalid!")
	}

}
func TestPositiveCaseMove(t *testing.T) {
	InputStdin := `
PLACE 0,0,NORTH
MOVE
MOVE
MOVE
`
	processLines(strings.NewReader(InputStdin))

	_, err := processCommands([]string{"MOVE"})
	fmt.Println(err)
	if err != nil {
		t.Error("Error position is invalid!")
	}

}
func TestNegativeCaseMove(t *testing.T) {
	InputStdin := `
PLACE 1,1,NORTH
MOVE
MOVE
MOVE
`
	processLines(strings.NewReader(InputStdin))

	_, err := processCommands([]string{"MOVE"})
	fmt.Println(err)
	if err == nil {
		t.Error("Error position is invalid!")
	}

}
