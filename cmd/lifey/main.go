package main

import (
	"github.com/ecshreve/lifey/internal/grid"
)

func main() {
	g := grid.NewGrid(3)
	g.Cells[0][0].Alive = true
	g.Cells[0][1].Alive = true
	g.Cells[0][2].Alive = true
	g.Cells[2][2].Alive = true
	g.PrintGrid()
	g.Tick()
	g.Tick()
	g.Tick()

	// g.PrintGrid()
	// for _, row := range g.Cells {
	// 	for _, c := range row {
	// 		fmt.Printf("cur: %v, next: %v -- r: %d, c: %d -- neighbors: %d\n", c.Alive, c.NextState, c.Row, c.Col, len(c.Neighbors))
	// 	}
	// }
	// fmt.Println("---")
	// g.SetNextState()

	// for _, row := range g.Cells {
	// 	for _, c := range row {
	// 		fmt.Printf("cur: %v, next: %v -- r: %d, c: %d -- neighbors: %d\n", c.Alive, c.NextState, c.Row, c.Col, len(c.Neighbors))
	// 	}
	// }
}
