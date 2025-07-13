package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pedro-git-projects/go-mtg/game"
)

// DrawIllustrations draws every entity that has both an Illustration and a Transform.
func DrawIllustrations(screen *ebiten.Image, g *game.Game) int {
	g.Lock.RLock()
	defer g.Lock.RUnlock()

	drawn := 0
	for idx, ill := range g.Illustrations {
		tx, ok := g.Transforms[idx]
		if !ok || ill.Image == nil {
			continue
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(tx.Scale, tx.Scale)
		op.GeoM.Translate(tx.X, tx.Y)
		screen.DrawImage(ill.Image, op)
		drawn++
	}
	return drawn
}
