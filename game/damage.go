package game

import "github.com/pedro-git-projects/go-mtg/component"

func (g *Game) DamageSystem(damage int) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	for idx, ent := range g.EM.Entites() {
		if ent.HasComponent(uint(component.LifeTotal)) {
			lt := g.lifeTotals[idx]
			lt.Value -= damage
			g.lifeTotals[idx] = lt
		}
	}
}
