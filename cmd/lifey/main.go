package main

import (
	"github.com/ecshreve/lifey/internal/grid"
)

func main() {
	g := grid.NewGrid(3)
	g.PrintGrid()
	g.PrintDebugInfo()

	g.Tick()
	g.PrintDebugInfo()
}
