package main

import (
	"github.com/ecshreve/lifey/internal/grid"
)

func main() {
	g := grid.NewGrid(5)
	g.Cells[0][0].Alive = true
	g.PrintGrid()
}
