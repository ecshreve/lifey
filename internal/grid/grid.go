package grid

import (
	"fmt"
	"math"
	"math/rand"
)

// Grid holds the data for the simulation.
type Grid struct {
	Size  int
	Cells [][]*Cell
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

// NewGrid returns a newly initialized Grid of Cells of the given size.
func NewGrid(size int) *Grid {
	cells := make([][]*Cell, size)
	for r := 0; r < size; r++ {
		row := make([]*Cell, size)
		for c := 0; c < size; c++ {
			row[c] = &Cell{
				Current:   Dead,
				Next:      Unknown,
				Row:       r,
				Col:       c,
				Neighbors: []*Cell{},
			}
		}
		cells[r] = row
	}
	g := &Grid{Size: size, Cells: cells}
	g.populateNeighbors()
	g.seedGrid(float64(0.75))
	return g
}

// seedGrid sets a percentage of the Grid's Cells to the Alive state based on
// the perc passed in.
func (g *Grid) seedGrid(perc float64) {
	initialAlive := int(math.Floor(float64(g.Size*g.Size) * perc))

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

// setNextState sets the NextState field for all Cells in the Grid.
func (g *Grid) setNextState() {
	for _, row := range g.Cells {
		for _, cell := range row {
			cell.setNextState()
		}
	}
}

// update sets the NextState field for all Cells in the Grid, and then calls the
// Cell.update() for each Cell in the Grid.
func (g *Grid) update() {
	g.setNextState()
	for _, row := range g.Cells {
		for _, cell := range row {
			cell.update()
		}
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

// PrintGrid prints a friendly representation of the Grid, where a Cell marked
// `A` represents an "Alive" Cell.
//
// For example a 3x3 Grid could look like this:
//
// +---+---+---+
// |   |   |   |
// +---+---+---+
// | A |   |   |
// +---+---+---+
// |   | A |   |
// +---+---+---+
func (g *Grid) PrintGrid() {
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
	fmt.Print(ret)
}

// PrintDebugInfo prints details about the current state of the Grid.
func (g *Grid) PrintDebugInfo() {
	for _, row := range g.Cells {
		for _, cell := range row {
			fmt.Println(cell)
		}
	}
}
