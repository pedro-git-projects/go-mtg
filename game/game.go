package game

import (
	"fmt"
	"sync"

	"github.com/pedro-git-projects/go-mtg/component"
	"github.com/pedro-git-projects/go-mtg/entity"
)

type Game struct {
	EM *entity.EntityManager

	lifeTotals      map[uint32]component.LifeTotalComponent
	names           map[uint32]component.NameComponent
	manaCosts       map[uint32]component.ManaCostComponent
	zoneTypes       map[uint32]component.ZoneComponent
	owners          map[uint32]component.OwnerComponent
	colorIndicators map[uint32]component.ColorIndicatorComponent
	contains        map[uint32]component.ContainsComponent

	Players []entity.Entity
	lock    sync.RWMutex
}

func NewGame(capacity int) *Game {
	return &Game{
		EM:              entity.NewEntityManager(capacity),
		lifeTotals:      make(map[uint32]component.LifeTotalComponent),
		names:           make(map[uint32]component.NameComponent),
		manaCosts:       make(map[uint32]component.ManaCostComponent),
		zoneTypes:       make(map[uint32]component.ZoneComponent),
		owners:          make(map[uint32]component.OwnerComponent),
		contains:        make(map[uint32]component.ContainsComponent),
		colorIndicators: make(map[uint32]component.ColorIndicatorComponent),
		Players:         make([]entity.Entity, 0, capacity),
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

func (g *Game) SpawnCard(cfg CardConfig) entity.Entity {
	g.lock.Lock()
	defer g.lock.Unlock()
	c := g.EM.Create()
	c.AddComponent(uint(component.Name))
	c.AddComponent(uint(component.ManaCost))
	c.AddComponent(uint(component.ColorIndicator))

	g.names[c.ID().Index] = component.NameComponent{Value: cfg.Name}
	g.manaCosts[c.ID().Index] = cfg.ManaCost
	g.colorIndicators[c.ID().Index] = cfg.ColorIndicator

	return c
}

func (g *Game) SpawnZone(zoneType component.Zone, owner entity.Entity, contents []entity.Entity) entity.Entity {
	g.lock.Lock()
	defer g.lock.Unlock()
	z := g.EM.Create()
	z.AddComponent(uint(component.ZoneType))
	z.AddComponent(uint(component.Owner))
	z.AddComponent(uint(component.Contains))

	g.zoneTypes[z.ID().Index] = component.ZoneComponent{Value: zoneType}
	g.owners[z.ID().Index] = component.OwnerComponent{Value: owner.ID()}

	ids := make([]entity.EntityId, len(contents))
	for i, e := range contents {
		ids[i] = e.ID()
	}
	g.contains[z.ID().Index] = component.ContainsComponent{Value: ids}

	return z
}

func (g *Game) SpawnPlayers(count int, initialLT int) {
	for i := 0; i < count; i++ {
		ent := g.SpawnWithLifeTotal(initialLT)
		g.lock.Lock()
		g.Players = append(g.Players, ent)
		g.lock.Unlock()
	}
}

func (g *Game) SpawnPlayer(initialLT int) entity.Entity {
	ent := g.SpawnWithLifeTotal(initialLT)
	g.lock.Lock()
	g.Players = append(g.Players, ent)
	g.lock.Unlock()
	return ent
}

func (g *Game) SpawnPlayersWithLibraries(count, initialLT int, cardConfigs [][]CardConfig) {
	if len(cardConfigs) != count {
		panic(fmt.Sprintf("expected %d slices of CardConfig, got %d", count, len(cardConfigs)))
	}
	for i := 0; i < count; i++ {
		player := g.SpawnPlayer(initialLT)

		var cards []entity.Entity
		for _, cfg := range cardConfigs[i] {
			cards = append(cards, g.SpawnCard(cfg))
		}

		g.SpawnZone(component.Library, player, cards)
	}
}
