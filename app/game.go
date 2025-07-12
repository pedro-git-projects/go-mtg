package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pedro-git-projects/go-mtg/game"
)

type App struct {
	Game *game.Game
}

func setupStandardGame() *game.Game {
	g := game.NewGame(10)
	g.SpawnPlayers(2, 20)
	return g
}

func NewAppp() *App {
	g := setupStandardGame()
	return &App{
		Game: g,
	}
}

func (app *App) Update() error {
	return nil
}

func (app *App) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, app.Game.PlayerStr())
}

func (app *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
