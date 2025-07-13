package app

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pedro-git-projects/go-mtg/component"
	"github.com/pedro-git-projects/go-mtg/game"
	"github.com/pedro-git-projects/go-mtg/system"
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
	if err := system.LoadIllustrations(g, "assets/cards"); err != nil {
		log.Fatalf("failed to load illustrations: %v", err)
	}

	for idx := range g.Illustrations {
		g.Transforms[idx] = component.TransformComponent{
			X:     float64((idx%5)*60 + 10), // or any debug position
			Y:     float64((idx/5)*90 + 10),
			Scale: 0.25, // try small so it fits your window
		}
	}
	return &App{Game: g}
}

func (app *App) Update() error {
	return nil
}

func (app *App) Draw(screen *ebiten.Image) {
	// clear to dark gray
	screen.Fill(color.RGBA{30, 30, 30, 255})

	total := len(app.Game.Illustrations)
	if total == 0 {
		ebitenutil.DebugPrint(screen, "No cards to draw")
		return
	}

	// compute grid layout
	cols := int(math.Ceil(math.Sqrt(float64(total))))
	bounds := screen.Bounds()
	sw, sh := bounds.Dx(), bounds.Dy() // replaced screen.Size()
	gutter := 10
	cellW := (sw - (cols+1)*gutter) / cols
	rows := int(math.Ceil(float64(total) / float64(cols)))
	cellH := (sh - (rows+1)*gutter) / rows

	i := 0
	for _, ill := range app.Game.Illustrations { // drop unused idx
		if ill.Image == nil {
			continue
		}
		row := i / cols
		col := i % cols

		// preserve aspect ratio
		iw := ill.Image.Bounds().Dx()
		ih := ill.Image.Bounds().Dy()
		scale := math.Min(
			float64(cellW)/float64(iw),
			float64(cellH)/float64(ih),
		)

		x := float64(gutter + col*(cellW+gutter))
		y := float64(gutter + row*(cellH+gutter))

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(x, y)
		screen.DrawImage(ill.Image, op)

		i++
	}

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("Cards: %d (%dx%d grid)", total, cols, rows),
		10, sh-20,
	)
}

func (app *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
