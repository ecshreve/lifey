package main

import (
	"fmt"

	"github.com/ecshreve/lifey/internal/grid"
)

func main() {
	g := grid.NewGrid(3)
	g.Cells[0][0].Alive = true

	g.PrintGrid()
	for _, row := range g.Cells {
		for _, c := range row {
			fmt.Printf("r: %d, c: %d -- neighbors: %d\n", c.Row, c.Col, len(c.Neighbors))
		}
	}
}
