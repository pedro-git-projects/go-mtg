package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pedro-git-projects/go-mtg/app"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Go MTG")
	app := app.NewApp()
	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
