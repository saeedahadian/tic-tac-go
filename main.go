package main

import (
	"log"
)

func main() {
	board := initGameBoard()
	board.Print()
}

func initGameBoard() *Board {
	board, err := NewBoard(3)
	if err != nil {
		log.Fatal(err)
	}
	return board
}
