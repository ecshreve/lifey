package grid

// Tick updates the Grid and prints the updated Grid.
func (g *Grid) Tick() {
	g.update()
	g.PrintGrid()
}
