
# TOY ROBOT SIMULATION

`Test coverage is: 90.3%`

## Pre-Requisites:
Go 1.10 installed with $GOPATH configured.

## Installation:
`go get github.com/alejoloaiza/toyrobotsimulator`

Go to the path: 
`cd $PATH/src/github.com/alejoloaiza/toyrobotsimulator`


##Run tests and show coverage:

`go test -cover`

Execute the program:
`go run main.go`

#Test data:

**Dataset 1:**
PLACE 1,1,NORTH
MOVE
MOVE
REPORT
**Expected output:**


**Dataset 2:**
PLACE 0,0,EAST
MOVE
**Expected output:**


