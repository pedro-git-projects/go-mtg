package app

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"sort"
	"strings"

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
	if err := system.LoadSymbols(g, "assets/mana"); err != nil {
		log.Fatalf("failed to load symbols: %v", err)
	}
	log.Printf("Loaded %d symbols:", len(g.Symbols))
	for k := range g.Symbols {
		log.Printf("  symbol key: %q", k)
	}
	for idx := range g.Illustrations {
		g.Transforms[idx] = component.TransformComponent{
			X:     float64((idx%5)*60 + 10),
			Y:     float64((idx/5)*90 + 10),
			Scale: 0.25,
		}
	}
	return &App{Game: g}
}

func (app *App) Update() error {
	return nil
}

func (app *App) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 30, 255})

	ids := sortedIllustrationIDs(app.Game)
	total := len(ids)
	if total == 0 {
		ebitenutil.DebugPrint(screen, "No cards to draw")
		return
	}

	cols, rows, cellW, cellH := computeGrid(screen, total)
	cardW := float64(cellW)
	cardH := cardW * 7.0 / 10.0
	artH := cardH * 0.6

	drawCount := 0
	for _, entID := range ids {
		ill := app.Game.Illustrations[entID]
		if ill.Image == nil {
			continue
		}

		x0, y0 := cellOrigin(drawCount, cols, cellW, cellH)
		drawCard(screen, app, entID, x0, y0, cardW, artH)

		drawCount++
		if drawCount >= total {
			break
		}
	}

	footer := fmt.Sprintf("Cards: %d (%dx%d)", drawCount, cols, rows)
	_, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	ebitenutil.DebugPrintAt(screen, footer, 10, sh-20)
}

func sortedIllustrationIDs(g *game.Game) []uint32 {
	m := g.Illustrations
	ids := make([]uint32, 0, len(m))
	for id := range m {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	return ids
}

func computeGrid(screen *ebiten.Image, total int) (cols, rows, cellW, cellH int) {
	cols = int(math.Ceil(math.Sqrt(float64(total))))
	bounds := screen.Bounds()
	sw, sh := bounds.Dx(), bounds.Dy()
	gutter := 10
	cellW = (sw - (cols+1)*gutter) / cols
	rows = int(math.Ceil(float64(total) / float64(cols)))
	cellH = (sh - (rows+1)*gutter) / rows
	return
}

func cellOrigin(index, cols, cellW, cellH int) (x0, y0 float64) {
	gutter := 10
	row := index / cols
	col := index % cols
	x0 = float64(gutter + col*(cellW+gutter))
	y0 = float64(gutter + row*(cellH+gutter))
	return
}

func drawCard(screen *ebiten.Image, app *App, entID uint32, x0, y0, cardW, artH float64) {
	drawArt(screen, app, entID, x0, y0, cardW, artH)
	drawName(screen, app, entID, x0, y0)
	drawSymbols(screen, app, entID, x0, y0, cardW)
	drawTypeLine(screen, app, entID, x0, y0, artH)
}

func drawArt(screen *ebiten.Image, app *App, entID uint32, x0, y0, cardW, artH float64) {
	img := app.Game.Illustrations[entID].Image
	iw, ih := img.Bounds().Dx(), img.Bounds().Dy()
	scale := math.Min(cardW/float64(iw), artH/float64(ih))
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x0, y0)
	screen.DrawImage(img, op)
}

func drawName(screen *ebiten.Image, app *App, entID uint32, x0, y0 float64) {
	name := app.Game.Names[entID].Value
	ebitenutil.DebugPrintAt(screen, name, int(x0)+8, int(y0)+8)
}

func drawSymbols(screen *ebiten.Image, app *App, entID uint32, x0, y0, cardW float64) {
	cost := app.Game.ManaCosts[entID]
	syms := collectSymbols(cost)
	symSize := 24.0
	for j := len(syms) - 1; j >= 0; j-- {
		key := syms[j]
		img := app.Game.Symbols[key]
		if img == nil {
			continue
		}
		scale := symSize / float64(img.Bounds().Dx())
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)
		x := x0 + cardW - 8 - symSize*float64(len(syms)-j)
		y := y0 + 8
		op.GeoM.Translate(x, y)
		screen.DrawImage(img, op)
	}
}

func collectSymbols(cost component.ManaCostComponent) []string {
	var s []string
	switch {
	case cost.Generic == 1:
		s = append(s, "One")
	case cost.Generic == 2:
		s = append(s, "Two")
	case cost.Generic == 3:
		s = append(s, "Three")
	case cost.Generic == 4:
		s = append(s, "Four")
	case cost.Generic > 4:
		for i := uint8(0); i < cost.Generic; i++ {
			s = append(s, "One")
		}
	}
	for i := uint8(0); i < cost.Colorless; i++ {
		s = append(s, "C")
	}
	for i := uint8(0); i < cost.White; i++ {
		s = append(s, "W")
	}
	for i := uint8(0); i < cost.Blue; i++ {
		s = append(s, "U")
	}
	for i := uint8(0); i < cost.Black; i++ {
		s = append(s, "B")
	}
	for i := uint8(0); i < cost.Red; i++ {
		s = append(s, "R")
	}
	for i := uint8(0); i < cost.Green; i++ {
		s = append(s, "G")
	}
	for i := uint8(0); i < cost.PhyWhite; i++ {
		s = append(s, "WP")
	}
	for i := uint8(0); i < cost.PhyBlue; i++ {
		s = append(s, "UP")
	}
	for i := uint8(0); i < cost.PhyBlack; i++ {
		s = append(s, "BP")
	}
	for i := uint8(0); i < cost.PhyRed; i++ {
		s = append(s, "RP")
	}
	for i := uint8(0); i < cost.PhyGreen; i++ {
		s = append(s, "GP")
	}
	for i := uint8(0); i < cost.HybridWU; i++ {
		s = append(s, "WU")
	}
	for i := uint8(0); i < cost.HybridWB; i++ {
		s = append(s, "WB")
	}
	for i := uint8(0); i < cost.HybridWG; i++ {
		s = append(s, "WG")
	}
	for i := uint8(0); i < cost.HybridWR; i++ {
		s = append(s, "WR")
	}
	for i := uint8(0); i < cost.HybridUB; i++ {
		s = append(s, "UB")
	}
	for i := uint8(0); i < cost.HybridUG; i++ {
		s = append(s, "UG")
	}
	for i := uint8(0); i < cost.HybridUR; i++ {
		s = append(s, "UR")
	}
	for i := uint8(0); i < cost.HybridBG; i++ {
		s = append(s, "BG")
	}
	for i := uint8(0); i < cost.HybridBR; i++ {
		s = append(s, "BR")
	}
	for i := uint8(0); i < cost.HybridGR; i++ {
		s = append(s, "GR")
	}
	return s
}

func drawTypeLine(screen *ebiten.Image, app *App, entID uint32, x0, y0, artH float64) {
	tl := app.Game.TypeLines[entID]
	var tparts, sparts []string
	for t := component.Artifact; t <= component.Vanguard; t++ {
		if tl.HasType(t) {
			tparts = append(tparts, t.String())
		}
	}
	for _, st := range tl.Subtypes {
		sparts = append(sparts, st.String())
	}
	line := strings.Join(tparts, " ")
	if len(sparts) > 0 {
		line += " — " + strings.Join(sparts, " ")
	}
	ebitenutil.DebugPrintAt(screen, line, int(x0)+8, int(y0+artH)+8)
}

func (app *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
