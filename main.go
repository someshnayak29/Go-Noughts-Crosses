package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var board [3][3]int
var turn int

const (
	NONE  = 0
	ROUND = 1
	CROSS = -1
)

type Location struct {
	x   int
	y   int
	val int
}

func main() {
	wg, locationCh, signalCh := InitGame()
	go NextMove(wg, locationCh, signalCh)
	go Game(wg, locationCh, signalCh)
	wg.Wait()
}

func InitGame() (*sync.WaitGroup, chan Location, chan bool) {
	var wg sync.WaitGroup
	wg.Add(2) // since we are starting 2 go routines
	locationCh := make(chan Location)
	signalCh := make(chan bool)
	turn = ROUND // game starts with Noughts

	return &wg, locationCh, signalCh
}

func Game(wg *sync.WaitGroup, locationCh chan Location, signalCh chan bool) {

	defer wg.Done() // decrease wait count on completion of this go routine
	for signalCh <- true; ; signalCh <- true {

		if ok := YourMove(locationCh); !ok {
			fmt.Println("Invalid Move, Try Again!")
			continue
		}

		turn = turn * -1 // Next turn
		PrintGame()

		if winner := GetWinner(); winner != NONE {
			fmt.Printf("Game over and the Winner is : %s \n", ConvertToChar(winner))
			signalCh <- false // End Game
			break
		}
	}
}

func NextMove(wg *sync.WaitGroup, locationCh chan Location, signalCh chan bool) {
	defer wg.Done()
	for {
		if sign := <-signalCh; !sign {
			break
		}

		x, y := GetCoordinates()
		locationCh <- Location{int(x), int(y), turn}
	}
}

func GetCoordinates() (int64, int64) {

	for {
		fmt.Println("x y >")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		s := sc.Text()

		if len(s) == 0 {
			fmt.Println("Invalid position to place stone")
			continue
		}

		cord := strings.Fields(s)
		if len(cord) != 2 {
			fmt.Println("Invalid position to place stone")
			continue
		}

		x, _ := strconv.ParseInt(cord[0], 10, 0)
		y, _ := strconv.ParseInt(cord[1], 10, 0)

		return x, y
	}
}

func PrintGame() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c := ConvertToChar(board[i][j])
			fmt.Printf("%s|", c)
		}
		fmt.Println()
	}
}

func ConvertToChar(cord int) string {

	if cord == NONE {
		return " "
	}

	if cord == ROUND {
		return "0"

	} else {
		return "X"
	}
}

func GetWinner() int {

	// Horizontal Match
	for i := 0; i < 3; i++ {
		if board[i][0] != NONE && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			return board[i][0]
		}
	}

	// Vertical Match
	for i := 0; i < 3; i++ {
		if board[0][i] != NONE && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			return board[0][i]
		}
	}

	// Both diagonals match( left diagonal, right diagonal)
	if board[0][0] != NONE && board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		return board[0][0]
	}
	if board[0][0] != NONE && board[0][2] == board[1][1] && board[0][2] == board[2][0] {
		return board[0][2]
	}
	return NONE

}

func YourMove(locationCh chan Location) bool {

	location := <-locationCh

	if location.x < 0 || location.x >= 3 || location.y < 0 || location.y >= 3 {
		return false
	}

	if board[location.x][location.y] != NONE {
		return false
	}

	board[location.x][location.y] = location.val
	return true
}
