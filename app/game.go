package app

import (
	"log"

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

func NewApp() *App {
	g, err := game.SetupStandardGameFromTOML("game.toml")
	if err != nil {
		log.Fatal(err)
	}
	return &App{Game: g}
}

func (app *App) Update() error {
	return nil
}

func (app *App) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, app.Game.LibraryStr())
}

func (app *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
