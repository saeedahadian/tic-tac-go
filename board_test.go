package main

import (
	"fmt"
	"testing"
)

func TestReturnErrorForBoardOfNegativeSize(t *testing.T) {
	_, err := NewBoard(-10)

	if err == nil {
		t.Errorf("Board of negative size should not be created")
	}
}

func TestCreateBoardWithPositiveSize(t *testing.T) {
	boardSize := 3
	board, err := NewBoard(boardSize)

	if err != nil {
		t.Errorf("Board should be created successfully with size of %d.", boardSize)
	}

	boardCellsCount := boardSize * boardSize
	if len(board.Cells) != boardCellsCount {
		t.Errorf("A board of size %d should include exactly %d cells but has %d cells.", boardSize, boardCellsCount, len(board.Cells))
	}
}

func TestInitializedCellHasCorrectIndices(t *testing.T) {
	boardSize := 3
	board, err := NewBoard(boardSize)
	if err != nil {
		t.Errorf("Board should be created successfully with the size of %d.", boardSize)
	}

	for i := range (boardSize * boardSize) {
		fmt.Printf("Here is the index: %d\n", i)
		cellIdx := board.Cells[i].Idx
		if cellIdx != i {
			t.Errorf("Cell at %dth index of the array should have the same index. Actual index: %d", i, cellIdx)
		}
	}
}

func FuzzInitializeBoardWithDifferentSizes(f *testing.F) {
	testcases := []int{-3, 0, 5}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, size int){
		_, err := NewBoard(size)
		if size > 0 && err != nil {
			t.Errorf("Board should be created successfully with given size: %d", size)
		}

		if size <= 0 && err == nil {
			t.Errorf("Board should have a positive size. Given %d", size)
		}
	})
}