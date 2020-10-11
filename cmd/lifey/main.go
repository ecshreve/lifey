package main

import (
	"github.com/ecshreve/lifey/internal/grid"
	"github.com/ecshreve/lifey/internal/sim"
)

func main() {
	g := grid.NewGrid(3)
	g.Tick()
	g.PrintDebugInfo()

	sim.StartSim()
}
