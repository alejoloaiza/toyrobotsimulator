package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	Place  string = "PLACE"
	Move   string = "MOVE"
	Left   string = "LEFT"
	Right  string = "RIGHT"
	Report string = "REPORT"

	North string = "NORTH"
	South string = "SOUTH"
	East  string = "EAST"
	West  string = "WEST"

	Started  string = "STARTED"
	Inactive string = "INACTIVE"
	Valid    string = "VALID"
	Invalid  string = "INVALID"

	CommandMaxNumFields = 2
	CommandMinNumFields = 1
	SizeOfMatrix        = 5
)

var (
	CurrentX           = 0
	CurrentY           = 0
	CurrentOrientation = ""
	CurrentStatus      = Inactive
)

func main() {
	processLines(os.Stdin)
}

func processLines(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		newline := scanner.Text()
		sline := strings.Split(newline, " ")
		if len(sline) <= CommandMaxNumFields || len(sline) >= CommandMinNumFields {
			processCommands(sline)
		} else {
			fmt.Println("Invalid command")
		}
	}
}

func processCommands(command []string) (bool, error) {
	var err error
	switch command[0] {
	case Place:
		err = place(command)
	case Move:
		_, err = move()
	case Right:
		right()
	case Left:
		left()
	case Report:
		report()
	default:
		fmt.Println("Invalid command!")
	}
	return false, err
}
func report() {
	if CurrentStatus == Started {
		fmt.Printf("Output: %v, %v, %s \n", CurrentX, CurrentY, CurrentOrientation)
	}
}
func place(command []string) error {
	if len(command) == 2 {
		sarg := strings.Split(command[1], ",")
		if len(sarg) == 3 {
			ProposedX, err := strconv.Atoi(sarg[0])
			if err != nil {
				return err
			}
			ProposedY, err := strconv.Atoi(sarg[1])
			if err != nil {
				return err
			}
			Result, err := validateProposal(ProposedX, ProposedY)
			if err != nil {
				//fmt.Println("Command ignored!")
				return err
			}
			if !validOrientation(sarg[2]) {
				return errors.New("Invalid orientation")
			}
			if Result {
				CurrentStatus = Started
				CurrentOrientation = sarg[2]
				CurrentX = ProposedX
				CurrentY = ProposedY
			}
		}

	}
	return nil
}

func validOrientation(name string) bool {
	if name == North || name == South || name == West || name == East {
		return true
	}
	return false
}

func right() (bool, error) {
	if CurrentStatus == Started {
		switch CurrentOrientation {
		case North:
			CurrentOrientation = East
		case East:
			CurrentOrientation = South
		case South:
			CurrentOrientation = West
		case West:
			CurrentOrientation = North
		}
		return true, nil
	}
	return false, errors.New("Robot toy not started, place the toy first")
}
func left() (bool, error) {
	if CurrentStatus == Started {
		switch CurrentOrientation {
		case North:
			CurrentOrientation = West
		case West:
			CurrentOrientation = South
		case South:
			CurrentOrientation = East
		case East:
			CurrentOrientation = North
		}
		return true, nil
	}
	return false, errors.New("Robot toy not started, place the toy first")
}

func move() (bool, error) {
	var ProposedX, ProposedY int
	if CurrentStatus == Started {
		ProposedX = CurrentX
		ProposedY = CurrentY
		switch CurrentOrientation {
		case North:
			ProposedY = CurrentY + 1
		case West:
			ProposedX = CurrentX - 1
		case South:
			ProposedY = CurrentY - 1
		case East:
			ProposedX = CurrentX + 1
		}
		Result, err := validateProposal(ProposedX, ProposedY)
		if err != nil {
			//fmt.Println("Command ignored!")
			return false, err
		}
		if Result {
			CurrentX = ProposedX
			CurrentY = ProposedY
			return true, nil
		}
		return false, nil
	}
	return false, errors.New("Robot toy not started, place the toy first")
}

func validateProposal(newX int, newY int) (bool, error) {
	var FromRoutine chan string
	FromRoutine = make(chan string)
	go validateNewPosition(newX, newY, FromRoutine)
	select {
	case msg := <-FromRoutine:
		if msg == Valid {
			return true, nil
		} else {
			return false, errors.New("Invalid position")
		}
	}
}
func validateNewPosition(newX int, newY int, mychan chan string) {
	defer func() {
		if r := recover(); r != nil {
			mychan <- Invalid
			return
		}
	}()
	Board := make([][]string, SizeOfMatrix)
	for x := 0; x < SizeOfMatrix; x++ {
		Board[x] = make([]string, SizeOfMatrix)
	}
	Board[newX][newY] = "X"
	mychan <- Valid
	return
}
