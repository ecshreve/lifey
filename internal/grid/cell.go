package grid

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

// setNextState sets the NextState field for the given Cell based on the
// following rules:
//
// 1. Any live cell with two or three live neighbours survives.
// 2. Any dead cell with exactly three live neighbours becomes a live cell.
// 3. All other live cells die in the next generation. Similarly, all other dead
//    cells stay dead.
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

// update sets the given Cell's Alive field to the value of NextState and resets
// NextState to false.
func (c *Cell) update() {
	c.Alive = c.NextState
	c.NextState = false
}
