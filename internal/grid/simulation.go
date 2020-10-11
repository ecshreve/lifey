package grid

// setNextState sets the Next field for all Cells in the Grid.
func (g *Grid) setNextState() {
	for _, row := range g.Cells {
		for _, cell := range row {
			cell.setNextState()
		}
	}
}

// update sets the Next field for all Cells in the Grid, and then calls the
// Cell.update() for each Cell in the Grid.
func (g *Grid) update() {
	g.setNextState()
	for _, row := range g.Cells {
		for _, cell := range row {
			cell.update()
		}
	}

	g.PrevAlive = g.CurrentAlive
	g.PrevDead = g.CurrentDead

	g.CurrentAlive = g.GetNumAlive()
	g.CurrentDead = g.Size*g.Size - g.CurrentAlive
	g.CurrentTick++
}

// Tick updates the Grid and prints the updated Grid.
func (g *Grid) Tick() {
	g.update()
}
