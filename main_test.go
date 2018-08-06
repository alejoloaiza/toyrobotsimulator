package main

import (
	"strings"
	"testing"
)

func TestInvalidCommand1(t *testing.T) {
	InputStdin := `
ALEJO LOAIZA COMMAND
`
	processLines(strings.NewReader(InputStdin))

	_, err := processCommands([]string{"MOVE"})
	if err == nil {
		t.Fail()
	}
}
func TestInvalidCommand2(t *testing.T) {
	_, err := processCommands([]string{"move"})
	if err != nil {
		t.Fail()
	}
}
func TestCaseReport(t *testing.T) {
	InputStdin := `
REPORT
PLACE 1,1,NORTH
REPORT
`
	processLines(strings.NewReader(InputStdin))
	_, err := processCommands([]string{"MOVE"})
	if err != nil {
		t.Fail()
	}
}
func TestPositiveCasePlace(t *testing.T) {
	InputStdin := `
LEFT
`
	processLines(strings.NewReader(InputStdin))
	_, err := processCommands([]string{"PLACE", "1,1,NORTH"})
	if err != nil {
		t.Fail()
	}

}
func TestNegativeCasePlace1(t *testing.T) {
	InputStdin := `
LEFT
`
	processLines(strings.NewReader(InputStdin))
	_, err := processCommands([]string{"PLACE", "5,4,NORTH"})
	if err == nil {
		t.Fail()
	}

}

func TestNegativeCasePlace2(t *testing.T) {
	InputStdin := `
LEFT
`
	processLines(strings.NewReader(InputStdin))
	_, err := processCommands([]string{"PLACE", "A,4,NORTH"})
	if err == nil {
		t.Fail()
	}

}

func TestNegativeCasePlace3(t *testing.T) {
	InputStdin := `
LEFT
`
	processLines(strings.NewReader(InputStdin))
	_, err := processCommands([]string{"PLACE", "3,B,NORTH"})
	if err == nil {
		t.Fail()
	}

}

func TestNegativeCasePlace4(t *testing.T) {
	InputStdin := `
LEFT
`
	processLines(strings.NewReader(InputStdin))
	_, err := processCommands([]string{"PLACE", "3,3,NORT"})
	if err == nil {
		t.Fail()
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
	if err == nil {
		t.Fail()
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
	if err != nil {
		t.Fail()
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
	if err == nil {
		t.Fail()
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
	if err != nil {
		t.Fail()
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
	if err != nil {
		t.Fail()
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
	if err == nil {
		t.Fail()
	}

}
