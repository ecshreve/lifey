package grid

import "fmt"

// Tick updates the Grid and prints the updated Grid.
func (g *Grid) Tick() {
	g.update()
	fmt.Println(g.GetGridString())
}
