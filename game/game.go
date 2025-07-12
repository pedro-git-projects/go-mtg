package game

import (
	"fmt"
	"strings"
	"sync"

	"github.com/pedro-git-projects/go-mtg/component"
	"github.com/pedro-git-projects/go-mtg/entity"
)

type Game struct {
	EM         *entity.EntityManager
	lifeTotals map[uint32]component.LifeTotalComponent
	Players    []entity.Entity
	lock       sync.RWMutex
}

func (g *Game) PlayerStr() string {
	g.lock.RLock()
	defer g.lock.RUnlock()

	var sb strings.Builder
	for idx, p := range g.Players {
		ltComp := g.lifeTotals[p.ID().Index]
		sb.WriteString(fmt.Sprintf(
			"Player %d (Entity%d): %d life\n",
			idx+1,
			p.ID().Index,
			ltComp.Value,
		))
	}
	return sb.String()
}

func NewGame(capacity int) *Game {
	return &Game{
		EM:         entity.NewEntityManager(capacity),
		lifeTotals: make(map[uint32]component.LifeTotalComponent, 2),
		Players:    make([]entity.Entity, 0, capacity),
	}
}

func (g *Game) SpawnWithLifeTotal(initialLT int) entity.Entity {
	ent := g.EM.Create()
	ent.AddComponent(uint(component.LifeTotal))

	g.lock.Lock()
	g.lifeTotals[ent.ID().Index] = component.LifeTotalComponent{Value: initialLT}
	g.lock.Unlock()

	return ent
}

func (g *Game) SpawnPlayers(count int, initialLT int) {
	for i := 0; i < count; i++ {
		ent := g.SpawnWithLifeTotal(initialLT)
		g.Players = append(g.Players, ent)
	}
}
