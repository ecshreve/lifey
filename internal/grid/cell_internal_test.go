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
