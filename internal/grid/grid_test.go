package grid_test

import (
	"testing"

	"github.com/ecshreve/lifey/internal/grid"
	"github.com/samsarahq/go/snapshotter"
)

func TestNewGrid(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	testGrid := grid.NewGrid(3)
	testGridStr := testGrid.GetGridString()
	snap.Snapshot("basic grid", testGridStr)

	testBigGrid := grid.NewGrid(15)
	testBigGridStr := testBigGrid.GetGridString()
	snap.Snapshot("big grid", testBigGridStr)
}
