package main

import "testing"

func TestFindCellXAndYWithIndex(t *testing.T) {
	idx := 5
	cell, _ := NewCell(idx)

	if cell.XPos() != 1 && cell.YPos() != 2 {
		t.Errorf("Cell with index of %d is in the row of 1 and the column of 2.", idx)
	}
}

func TestReturnErrorOnceCellIndexIsOutOfRange(t *testing.T) {
	idx := 100

	_, err := NewCell(idx)
	if err == nil {
		t.Errorf("Cell index of 100 is out range of the board game.")
	}
}

func TestReturnErrorOnceCellIndexIsBelowZero(t *testing.T) {
	idx := -5

	_, err := NewCell(idx)
	if err == nil {
		t.Errorf("Index %d is lower than zero.", idx)
	}
}
