package sim

import (
	"os"
	"time"

	"github.com/gdamore/tcell"

	"github.com/ecshreve/lifey/internal/grid"
	"github.com/rivo/tview"
)

var (
	app        *tview.Application
	headerFlex *tview.Flex
	simFlex    *tview.Flex
	endButton  *tview.Button
	tickButton *tview.Button
)

var hasUpdate = make(chan bool)

func update(g *grid.Grid) {
	for {
		<-hasUpdate
		app.QueueUpdateDraw(func() {
			simFlex.Clear()
			for r := 0; r < g.Size; r++ {
				rowFlex := tview.NewFlex().SetDirection(tview.FlexColumn)
				for c := 0; c < g.Size; c++ {
					cell := tview.NewBox()

					if g.Cells[r][c].Current == grid.Alive {
						cell.SetBackgroundColor(tcell.ColorGreen)
					} else {
						cell.SetBackgroundColor(tcell.ColorDarkRed)
					}
					rowFlex.AddItem(cell, 0, 1, false)
				}
				simFlex.AddItem(rowFlex, 0, 1, false)
			}
			time.Sleep(time.Second / 4)
			app.SetFocus(tickButton)
		})
	}
}

func StartSim() {
	g := grid.NewGrid(10)
	app = tview.NewApplication()

	headerFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	simFlex = tview.NewFlex().SetDirection(tview.FlexRow)

	headerFlex.SetBorder(true).SetTitle(" header ").SetBorderPadding(1, 1, 1, 1)
	simFlex.SetBorder(true).SetTitle(" sim ").SetBorderPadding(1, 1, 1, 1)

	tickButton = tview.NewButton("tick")
	tickButton.SetBackgroundColorActivated(tcell.ColorSalmon)
	tickButton.SetBorder(true)
	tickButton.SetSelectedFunc(func() {
		g.Tick()
		app.SetFocus(simFlex)
		hasUpdate <- true
	})
	headerFlex.AddItem(tickButton, 0, 1, false)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(headerFlex, 0, 1, false).
			AddItem(simFlex, 0, 2, false), 0, 5, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune {
			switch event.Rune() {
			case 'q':
				app.Stop()
				os.Exit(0)
			}
		}
		return event
	})
	go update(g)
	if err := app.SetRoot(flex, true).SetFocus(tickButton).Run(); err != nil {
		panic(err)
	}

}
