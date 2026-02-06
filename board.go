package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	XUnicode = "\u274c"
	OUnicode = "\u2b55"
	EmptyUnicode = "\u26aa"
)

var (
	BoardSizeIsNegative = errors.New("Board size is negative")
)

type Board struct {
	Size  int
	Cells []*Cell
}

func (b *Board) Print() {
	var output []string
	for i := range len(b.Cells) {
		if i > 0 && (i % b.Size == 0)  {
			output = append(output, "\n")
		}

		if b.Cells[i].Val == nil {
			output = append(output, EmptyUnicode)
			continue
		}

		switch *b.Cells[i].Val {
		case X:
			output = append(output, XUnicode)
		case O:
			output = append(output, OUnicode)
		}
	}
	
	fmt.Println(strings.Join(output, ""))
}

func NewBoard(size int) (*Board, error) {
	if size <= 0 {
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
