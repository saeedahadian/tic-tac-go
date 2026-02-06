package main

import (
	"errors"
)

var (
	BoardSizeIsNegative = errors.New("Board size is negative")
)

type Board struct {
	Size  int
	Cells []*Cell
}

func NewBoard(size int) (*Board, error) {
	if size < 0 {
		return nil, BoardSizeIsNegative
	}

	cells := newBoardCells(size)

	return &Board{
		Size:  size,
		Cells: cells,
	}, nil
}

func newBoardCells(boardSize int) []*Cell {
	cellsCount := boardSize * boardSize
	cells := make([]*Cell, cellsCount)
	for i := range cellsCount {
		cells[i], _ = NewCell(i)
	}
	return cells
}
