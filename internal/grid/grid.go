package grid

import "fmt"

// Cell represents an individual cell in the simulation.
type Cell struct {
	Alive bool
	Row   int
	Col   int
}

// cellStr returns the string representation of the given Cell. " A " if the
// Cell is Alive, or "   " if it isn't.
func (c *Cell) cellStr() string {
	ret := " "

	if c.Alive {
		ret += "A"
	} else {
		ret += " "
	}

	ret += " "
	return ret
}

// Grid holds the data for the simulation.
type Grid struct {
	Size  int
	Cells [][]*Cell
}

// NewGrid returns a newly initialized Grid of Cells of the given size.
func NewGrid(size int) *Grid {
	cells := make([][]*Cell, size)
	for r := 0; r < size; r++ {
		row := make([]*Cell, size)
		for c := 0; c < size; c++ {
			row[c] = &Cell{
				Alive: false,
				Row:   r,
				Col:   c,
			}
		}
		cells[r] = row
	}
	return &Grid{Size: size, Cells: cells}
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
			rowStr += cell.cellStr()
			rowStr += "|"
		}
		rowStr += "\n"
		ret += rowStr
		ret += divider
	}
	fmt.Print(ret)
}
