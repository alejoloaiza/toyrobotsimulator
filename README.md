
# TOY ROBOT SIMULATION

`Test coverage is: 90.3%`

## Pre-Requisites:
Go 1.10 installed with $GOPATH configured.

## Installation:
`go get github.com/alejoloaiza/toyrobotsimulator`

Go to the path: 
`cd $GOPATH/src/github.com/alejoloaiza/toyrobotsimulator`


## Run tests and show coverage:

`go test -cover`

## Execute the program and run test data:
`go run main.go`

# Test data:

**Dataset 1:**      
PLACE 1,1,NORTH     
MOVE        
MOVE        
REPORT      
**Expected output:**
Output: 1, 3, NORTH

**Dataset 2:**      
PLACE 0,0,EAST      
MOVE        
MOVE        
MOVE        
MOVE        
REPORT      
**Expected output:**
Output: 4, 0, EAST

**Dataset 3:**
PLACE 3,0,EAST      
MOVE        
MOVE        
MOVE        
MOVE        
REPORT      
MOVE       
REPORT      
**Expected output:**
Output: 4, 0, EAST 

**Dataset 4:**
PLACE 3,0,WEST      
MOVE        
MOVE        
REPORT      
**Expected output:**
Output: 1, 0, WEST

**Dataset 5:**
PLACE 3,0,WEST      
MOVE        
MOVE        
LEFT        
REPORT      
**Expected output:**
Output: 1, 0, SOUTH

**Dataset 6:**
PLACE 0,0,NORTH     
LEFT        
LEFT        
LEFT        
LEFT        
REPORT      
**Expected output:**
Output: 0, 0, NORTH

**Dataset 7:**
PLACE 0,0,NORTH     
RIGHT       
RIGHT       
RIGHT       
REPORT      
**Expected output:**
Output: 0, 0, WEST

**Dataset 8:**
PLACE 0,0,NORTH     
MOVE        
RIGHT       
MOVE        
RIGHT       
MOVE        
RIGHT       
MOVE        
RIGHT       
REPORT      
**Expected output:**
Output: 0, 0, NORTH