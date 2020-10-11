package grid

import (
	"fmt"
	"math"
	"math/rand"
)

// Grid holds the data for the simulation.
type Grid struct {
	Size         int
	Cells        [][]*Cell
	InitialAlive int
	InitialDead  int
	PrevAlive    int
	PrevDead     int
	CurrentAlive int
	CurrentDead  int
	CurrentTick  int
}

// NewGrid returns a newly initialized Grid of Cells of the given size.
func NewGrid(size int) *Grid {
	cells := make([][]*Cell, size)
	for r := 0; r < size; r++ {
		row := make([]*Cell, size)
		for c := 0; c < size; c++ {
			row[c] = &Cell{
				Current:   Dead,
				Row:       r,
				Col:       c,
				Neighbors: []*Cell{},
			}
		}
		cells[r] = row
	}
	g := &Grid{Size: size, Cells: cells}
	g.populateNeighbors()
	g.seedGrid(float64(0.4))
	return g
}

// populateNeighbors sets the Cell.Neighbors value for each Cell in the Grid.
func (g *Grid) populateNeighbors() {
	for _, row := range g.Cells {
		for _, cell := range row {
			neighborIndices := cell.getNeighborIndices(g.Size)
			for _, neighborIndex := range neighborIndices {
				cell.Neighbors = append(cell.Neighbors, g.Cells[neighborIndex[0]][neighborIndex[1]])
			}
		}
	}
}

// seedGrid sets a percentage of the Grid's Cells to the Alive state based on
// the perc passed in.
func (g *Grid) seedGrid(perc float64) {
	initialAlive := int(math.Floor(float64(g.Size*g.Size) * perc))
	g.InitialAlive = initialAlive
	g.InitialDead = g.Size*g.Size - initialAlive
	g.CurrentAlive = g.InitialAlive
	g.CurrentDead = g.InitialDead

	// TODO: change this seed value to be based on current time so it's really
	// random, leaving as a hard-coded value to make testing easier.
	rand.Seed(3)
	for initialAlive > 0 {
		// Pick a random row and column value.
		r := rand.Intn(g.Size)
		c := rand.Intn(g.Size)

		// If the Cell at that row and column is already Alive then skip to the
		// next iteration and pick a new row and column.
		if g.Cells[r][c].Current == Alive {
			continue
		}

		g.Cells[r][c].Current = Alive
		initialAlive--
	}
}

// getDivider returns the divider string based on the Grid's size.
func (g *Grid) getDivider() string {
	ret := "+"
	for i := 0; i < g.Size; i++ {
		ret += "---+"
	}
	ret += "\n"
	return ret
}

// GetNumAlive returns the total number of currently Alive Cells in the Grid.
func (g *Grid) GetNumAlive() int {
	numAlive := 0
	for _, row := range g.Cells {
		numAlive += getNumAlive(row)
	}
	return numAlive
}

// GetGridString returns a friendly representation of the Grid, where a Cell
// marked `A` represents an "Alive" Cell, and one marked "." is Dead.
//
// For example a 3x3 Grid could look like this:
//
// +---+---+---+
// | . | A | . |
// +---+---+---+
// | A | . | . |
// +---+---+---+
// | . | A | . |
// +---+---+---+
func (g *Grid) GetGridString() string {
	divider := g.getDivider()
	ret := ""
	ret += divider
	for _, row := range g.Cells {
		rowStr := "|"
		for _, cell := range row {
			rowStr += fmt.Sprintf(" %v |", cell.Current)
		}
		rowStr += "\n"
		ret += rowStr
		ret += divider
	}
	return ret
}

// PrintDebugInfo prints details about the current state of the Grid.
func (g *Grid) PrintDebugInfo() {
	for _, row := range g.Cells {
		for _, cell := range row {
			fmt.Println(cell)
		}
	}
}
