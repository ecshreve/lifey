package sim

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"

	"github.com/ecshreve/lifey/internal/grid"
	"github.com/rivo/tview"
)

var (
	app        *tview.Application
	infoFlex   *tview.Flex
	statFlex   *tview.Flex
	simFlex    *tview.Flex
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

			statFlex.Clear()
			simStatStr := fmt.Sprintf("current tick: %d\n---\nsize: %d\n---\nalive: %d\ndead: %d\n", g.CurrentTick, g.Size, g.InitialAlive, g.InitialDead)
			simStatText := tview.NewTextView().SetText(simStatStr).SetTextAlign(0).SetWordWrap(true)
			simStatText.SetBorder(true).SetBorderPadding(1, 0, 2, 1).SetTitle(" overall ")

			prevStatStr := fmt.Sprintf("---\nalive: %d\ndead: %d\n---", g.PrevAlive, g.PrevDead)
			prevStatText := tview.NewTextView().SetText(prevStatStr).SetTextAlign(0).SetWordWrap(true)
			prevStatText.SetBorder(true).SetBorderPadding(1, 0, 2, 1).SetTitle(" previous ")

			curStatStr := fmt.Sprintf("---\nalive: %d\ndead: %d\n---", g.CurrentAlive, g.CurrentDead)
			curStatText := tview.NewTextView().SetText(curStatStr).SetTextAlign(0).SetWordWrap(true)
			curStatText.SetBorder(true).SetBorderPadding(1, 0, 2, 1).SetTitle(" current ")

			statFlex.AddItem(simStatText, 0, 1, false)
			statFlex.AddItem(prevStatText, 0, 1, false)
			statFlex.AddItem(curStatText, 0, 1, false)

			time.Sleep(time.Second / 4)
			app.SetFocus(tickButton)
		})
	}
}

func StartSim() {
	g := grid.NewGrid(100)
	app = tview.NewApplication()

	infoFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	statFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	simFlex = tview.NewFlex().SetDirection(tview.FlexRow)

	infoFlex.SetBorder(true).SetTitle(" info ").SetBorderPadding(1, 1, 1, 1)
	statFlex.SetBorder(true).SetTitle(" stats ").SetBorderPadding(1, 1, 1, 1)
	simFlex.SetBorder(true).SetTitle(" sim ").SetBorderPadding(1, 1, 1, 1)

	instructions := "press `q` to exit\n\npress `enter` to trigger the simulation's Tick() function"
	instructionsText := tview.NewTextView().SetText(instructions)
	instructionsText.SetBorderPadding(2, 1, 2, 1)
	instructionsText.SetWrap(true).SetWordWrap(true)
	infoFlex.AddItem(instructionsText, 0, 1, false)

	tickButton = tview.NewButton("tick")
	tickButton.SetBackgroundColorActivated(tcell.ColorSalmon)
	tickButton.SetBorder(true)
	tickButton.SetSelectedFunc(func() {
		g.Tick()
		app.SetFocus(simFlex)
		hasUpdate <- true
	})
	infoFlex.AddItem(tickButton, 0, 1, false)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(infoFlex, 0, 1, false).
			AddItem(statFlex, 0, 1, false).
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
	hasUpdate <- true

	if err := app.SetRoot(flex, true).SetFocus(tickButton).Run(); err != nil {
		panic(err)
	}

}
