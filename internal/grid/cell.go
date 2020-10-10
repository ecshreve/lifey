package grid

// CellState represents the state of a Cell.
type CellState int

// Enum defining the possible CellStates.
const (
	Dead CellState = iota
	Alive
	Unknown
)

func (cs CellState) String() string {
	return [...]string{" ", "A", "u"}[cs]
}

// Cell represents an individual cell in the simulation.
type Cell struct {
	Current   CellState
	Next      CellState
	Row       int
	Col       int
	Neighbors []*Cell
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
		if c.Current == Alive {
			numAlive++
		}
	}
	return numAlive
}

// setNextState sets the Next field for the given Cell based on the
// following rules:
//
// 1. Any live cell with two or three live neighbours survives.
// 2. Any dead cell with exactly three live neighbours becomes a live cell.
// 3. All other live cells die in the next generation. Similarly, all other dead
//    cells stay dead.
func (c *Cell) setNextState() {
	numAlive := getNumAlive(c.Neighbors)

	nextState := Dead
	if c.Current == Alive {
		if numAlive == 2 || numAlive == 3 {
			nextState = Alive
		}
	} else {
		if numAlive == 3 {
			nextState = Alive
		}
	}

	c.Next = nextState
}

// update sets the given Cell's Current field to the value of Next and sets
// Next to Unknown.
func (c *Cell) update() {
	c.Current = c.Next
	c.Next = Unknown
}
