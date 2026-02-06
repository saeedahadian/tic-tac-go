package main

import (
	"fmt"
	"log"
)

func main() {
	board := initGameBoard()
	fmt.Printf("%v\n", board)
}

func initGameBoard() *Board {
	board, err := NewBoard(3)
	if err != nil {
		log.Fatal(err)
	}
	return board
}
