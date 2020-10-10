package grid

import "fmt"

// Cell represents an individual cell in the simulation.
type Cell struct {
	Alive     bool
	NextState bool
	Row       int
	Col       int
	Neighbors []*Cell
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

// getNeighborIndices returns a 2d slice of ints where each element is of the
// form []int{row, col}. These represent grid locations that are valid neighbors
// of the given Cell.
func (c *Cell) getNeighborIndices(size int) [][]int {
	possibles := [][]int{
		[]int{c.Row - 1, c.Col - 1},
		[]int{c.Row - 1, c.Col},
		[]int{c.Row - 1, c.Col + 1},

		[]int{c.Row, c.Col - 1},
		[]int{c.Row, c.Col + 1},

		[]int{c.Row + 1, c.Col - 1},
		[]int{c.Row + 1, c.Col},
		[]int{c.Row + 1, c.Col + 1},
	}

	var filtered [][]int
	for _, possible := range possibles {
		if possible[0] < 0 || possible[1] < 0 {
			continue
		}
		if possible[0] >= size || possible[1] >= size {
			continue
		}
		filtered = append(filtered, possible)
	}
	return filtered
}

// getNumAlive returns the number of Alive Cells in the given slice.
func getNumAlive(cells []*Cell) int {
	numAlive := 0
	for _, c := range cells {
		if c.Alive {
			numAlive++
		}
	}
	return numAlive
}

func (c *Cell) setNextState() {
	numAlive := getNumAlive(c.Neighbors)

	nextState := false
	if c.Alive {
		if numAlive == 2 || numAlive == 3 {
			nextState = true
		}
	} else {
		if numAlive == 3 {
			nextState = true
		}
	}

	c.NextState = nextState
}

func (c *Cell) updateState() {
	c.Alive = c.NextState
	c.NextState = false
}

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
				Alive:     false,
				Row:       r,
				Col:       c,
				Neighbors: []*Cell{},
			}
		}
		cells[r] = row
	}
	g := &Grid{Size: size, Cells: cells}
	g.populateNeighbors()
	return g
}

func (g *Grid) setNextState() {
	for _, row := range g.Cells {
		for _, cell := range row {
			cell.setNextState()
		}
	}
}

func (g *Grid) update() {
	g.setNextState()
	for _, row := range g.Cells {
		for _, cell := range row {
			cell.updateState()
		}
	}
}

func (g *Grid) Tick() {
	g.update()
	g.PrintGrid()
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
