package system

import (
	"github.com/pedro-git-projects/go-mtg/component"
	"github.com/pedro-git-projects/go-mtg/game"
)

func DamageSystem(g *game.Game, damage int) {
	g.Lock.RLock()
	defer g.Lock.RUnlock()

	for idx, ent := range g.EM.Entites() {
		if ent.HasComponent(uint(component.LifeTotal)) {
			lt := g.LifeTotals[idx]
			lt.Value -= damage
			g.LifeTotals[idx] = lt
		}
	}
}
