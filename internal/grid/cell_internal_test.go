package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNeighborIndices(t *testing.T) {
	testcases := []struct {
		desc     string
		cell     *Cell
		size     int
		expected [][]int
	}{
		{
			desc: "top left corner, 3 neighbors",
			cell: &Cell{
				Row: 0,
				Col: 0,
			},
			size: 3,
			expected: [][]int{
				[]int{0, 1},
				[]int{1, 1},
				[]int{1, 0},
			},
		},
		{
			desc: "top row, 5 neighbors",
			cell: &Cell{
				Row: 0,
				Col: 1,
			},
			size: 3,
			expected: [][]int{
				[]int{0, 0},
				[]int{0, 2},
				[]int{1, 0},
				[]int{1, 1},
				[]int{1, 2},
			},
		},
		{
			desc: "top right corner, 3 neighbors",
			cell: &Cell{
				Row: 0,
				Col: 2,
			},
			size: 3,
			expected: [][]int{
				[]int{0, 1},
				[]int{1, 1},
				[]int{1, 2},
			},
		},
		{
			desc: "left col, 5 neighbors",
			cell: &Cell{
				Row: 1,
				Col: 0,
			},
			size: 3,
			expected: [][]int{
				[]int{0, 0},
				[]int{0, 1},
				[]int{1, 1},
				[]int{2, 0},
				[]int{2, 1},
			},
		},
		{
			desc: "center, 8 neighbors",
			cell: &Cell{
				Row: 1,
				Col: 1,
			},
			size: 3,
			expected: [][]int{
				[]int{0, 0},
				[]int{0, 1},
				[]int{0, 2},
				[]int{1, 0},
				[]int{1, 2},
				[]int{2, 0},
				[]int{2, 1},
				[]int{2, 2},
			},
		},
		{
			desc: "right col, 5 neighbors",
			cell: &Cell{
				Row: 1,
				Col: 2,
			},
			size: 3,
			expected: [][]int{
				[]int{0, 1},
				[]int{0, 2},
				[]int{1, 1},
				[]int{2, 1},
				[]int{2, 2},
			},
		},
		{
			desc: "bottom left corner, 3 neighbors",
			cell: &Cell{
				Row: 2,
				Col: 0,
			},
			size: 3,
			expected: [][]int{
				[]int{1, 0},
				[]int{1, 1},
				[]int{2, 1},
			},
		},
		{
			desc: "bottom row, 5 neighbors",
			cell: &Cell{
				Row: 2,
				Col: 1,
			},
			size: 3,
			expected: [][]int{
				[]int{1, 0},
				[]int{1, 1},
				[]int{1, 2},
				[]int{2, 0},
				[]int{2, 2},
			},
		},
		{
			desc: "bottom right corner, 3 neighbors",
			cell: &Cell{
				Row: 2,
				Col: 2,
			},
			size: 3,
			expected: [][]int{
				[]int{1, 2},
				[]int{1, 1},
				[]int{2, 1},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual := testcase.cell.getNeighborIndices(testcase.size)
			assert.ElementsMatch(t, testcase.expected, actual)
		})
	}
}

func TestGetNumAlive(t *testing.T) {
	testcases := []struct {
		desc     string
		cells    []*Cell
		expected int
	}{
		{
			desc:     "empty slice returns zero",
			cells:    []*Cell{},
			expected: 0,
		},
		{
			desc: "slice with none alive returns zero",
			cells: []*Cell{
				&Cell{Current: Dead},
				&Cell{Current: Dead},
				&Cell{Current: Dead},
			},
			expected: 0,
		},
		{
			desc: "slice with one alive returns one",
			cells: []*Cell{
				&Cell{Current: Dead},
				&Cell{Current: Alive},
				&Cell{Current: Dead},
			},
			expected: 1,
		},
		{
			desc: "slice with multiple alive returns expected number",
			cells: []*Cell{
				&Cell{Current: Alive},
				&Cell{Current: Alive},
				&Cell{Current: Dead},
			},
			expected: 2,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual := getNumAlive(testcase.cells)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}

func TestSetNextState(t *testing.T) {
	testcases := []struct {
		desc     string
		cell     *Cell
		expected CellState
	}{
		{
			desc: "live cell with 0 live neighbor dies",
			cell: &Cell{
				Current: Alive,
				Neighbors: []*Cell{
					&Cell{Current: Dead},
					&Cell{Current: Dead},
					&Cell{Current: Dead},
					&Cell{Current: Dead},
				},
			},
			expected: Dead,
		},
		{
			desc: "live cell with 1 live neighbor dies",
			cell: &Cell{
				Current: Alive,
				Neighbors: []*Cell{
					&Cell{Current: Alive},
					&Cell{Current: Dead},
					&Cell{Current: Dead},
					&Cell{Current: Dead},
				},
			},
			expected: Dead,
		},
		{
			desc: "live cell with 2 live neighbors lives",
			cell: &Cell{
				Current: Alive,
				Neighbors: []*Cell{
					&Cell{Current: Alive},
					&Cell{Current: Alive},
					&Cell{Current: Dead},
					&Cell{Current: Dead},
				},
			},
			expected: Alive,
		},
		{
			desc: "live cell with 3 live neighbors lives",
			cell: &Cell{
				Current: Alive,
				Neighbors: []*Cell{
					&Cell{Current: Alive},
					&Cell{Current: Alive},
					&Cell{Current: Alive},
					&Cell{Current: Dead},
				},
			},
			expected: Alive,
		},
		{
			desc: "live cell with 4 live neighbors dies",
			cell: &Cell{
				Current: Alive,
				Neighbors: []*Cell{
					&Cell{Current: Alive},
					&Cell{Current: Alive},
					&Cell{Current: Alive},
					&Cell{Current: Alive},
				},
			},
			expected: Dead,
		},
		{
			desc: "dead cell with 0 live neighbor stays dead",
			cell: &Cell{
				Current: Dead,
				Neighbors: []*Cell{
					&Cell{Current: Dead},
					&Cell{Current: Dead},
					&Cell{Current: Dead},
					&Cell{Current: Dead},
				},
			},
			expected: Dead,
		},
		{
			desc: "dead cell with 1 live neighbor stays dead",
			cell: &Cell{
				Current: Dead,
				Neighbors: []*Cell{
					&Cell{Current: Alive},
					&Cell{Current: Dead},
					&Cell{Current: Dead},
					&Cell{Current: Dead},
				},
			},
			expected: Dead,
		},
		{
			desc: "dead cell with 2 live neighbors stays dead",
			cell: &Cell{
				Current: Dead,
				Neighbors: []*Cell{
					&Cell{Current: Alive},
					&Cell{Current: Alive},
					&Cell{Current: Dead},
					&Cell{Current: Dead},
				},
			},
			expected: Dead,
		},
		{
			desc: "dead cell with 3 live neighbors lives",
			cell: &Cell{
				Current: Dead,
				Neighbors: []*Cell{
					&Cell{Current: Alive},
					&Cell{Current: Alive},
					&Cell{Current: Alive},
					&Cell{Current: Dead},
				},
			},
			expected: Alive,
		},
		{
			desc: "dead cell with 4 live neighbors stays dead",
			cell: &Cell{
				Current: Dead,
				Neighbors: []*Cell{
					&Cell{Current: Alive},
					&Cell{Current: Alive},
					&Cell{Current: Alive},
					&Cell{Current: Alive},
				},
			},
			expected: Dead,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			testcase.cell.setNextState()
			assert.Equal(t, testcase.expected, testcase.cell.Next)
		})
	}
}
