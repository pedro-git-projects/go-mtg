package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type App struct {
}

func NewAppp() *App {
	return &App{}
}

func (app *App) Update() error {
	return nil
}

func (app *App) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "MTG Sucks")
}

func (app *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
