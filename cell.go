package main

import "fmt"

type Mark int

func (m *Mark) Print() {
	if m == nil {
		fmt.Println("nil")
		return
	}

	fmt.Printf("%d\n", m)
}

const (
	X Mark = iota
	O
	Empty
)

type Cell struct {
	Idx int
	Val *Mark
}

func (c *Cell) XPos() int {
	return c.Idx / 3
}

func (c *Cell) YPos() int {
	return c.Idx % 3
}

func (c *Cell) isMarked() bool {
	return c.Val != nil
}

func (c *Cell) isEmpty() bool {
	return c.Val == nil
}

func (c *Cell) isX() bool {
	return *c.Val == X
}

func (c *Cell) isO() bool {
	return *c.Val == O
}

func NewCell(idx int) (*Cell, error) {
	return &Cell{
		Idx: idx,
		Val: nil,
	}, nil
}
