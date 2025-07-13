package game

import (
	"fmt"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pedro-git-projects/go-mtg/component"
	"github.com/pedro-git-projects/go-mtg/entity"
)

type Game struct {
	EM *entity.EntityManager

	LifeTotals      map[uint32]component.LifeTotalComponent
	Names           map[uint32]component.NameComponent
	ManaCosts       map[uint32]component.ManaCostComponent
	TypeLines       map[uint32]component.TypeLineComponent
	ZoneTypes       map[uint32]component.ZoneComponent
	Owners          map[uint32]component.OwnerComponent
	ColorIndicators map[uint32]component.ColorIndicatorComponent
	Contains        map[uint32]component.ContainsComponent
	Illustrations   map[uint32]component.IllustrationComponent
	Transforms      map[uint32]component.TransformComponent
	Symbols         map[string]*ebiten.Image

	Players []entity.Entity
	Lock    sync.RWMutex
}

func NewGame(capacity int) *Game {
	return &Game{
		EM:              entity.NewEntityManager(capacity),
		LifeTotals:      make(map[uint32]component.LifeTotalComponent),
		Names:           make(map[uint32]component.NameComponent),
		ManaCosts:       make(map[uint32]component.ManaCostComponent),
		TypeLines:       make(map[uint32]component.TypeLineComponent),
		ZoneTypes:       make(map[uint32]component.ZoneComponent),
		Owners:          make(map[uint32]component.OwnerComponent),
		Contains:        make(map[uint32]component.ContainsComponent),
		ColorIndicators: make(map[uint32]component.ColorIndicatorComponent),
		Illustrations:   make(map[uint32]component.IllustrationComponent),
		Transforms:      make(map[uint32]component.TransformComponent),
		Players:         make([]entity.Entity, 0, capacity),
	}
}

func (g *Game) SpawnWithLifeTotal(initialLT int) entity.Entity {
	ent := g.EM.Create()
	ent.AddComponent(uint(component.LifeTotal))

	g.Lock.Lock()
	g.LifeTotals[ent.ID().Index] = component.LifeTotalComponent{Value: initialLT}
	g.Lock.Unlock()

	return ent
}

func (g *Game) SpawnCard(cfg CardConfig) entity.Entity {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	c := g.EM.Create()
	c.AddComponent(uint(component.Name))
	c.AddComponent(uint(component.ManaCost))
	c.AddComponent(uint(component.ColorIndicator))
	c.AddComponent(uint(component.TypeLine))

	g.Names[c.ID().Index] = component.NameComponent{Value: cfg.Name}
	g.ManaCosts[c.ID().Index] = cfg.ManaCost
	g.TypeLines[c.ID().Index] = cfg.TypeLine.ToComponent()
	g.ColorIndicators[c.ID().Index] = cfg.ColorIndicator

	if cfg.Illustration != nil {
		c.AddComponent(uint(component.Illustration))
		g.Illustrations[c.ID().Index] = component.IllustrationComponent{
			AssetKey: cfg.Illustration.AssetKey,
			Image:    nil,
		}
	}

	return c
}

func (g *Game) SpawnZone(zoneType component.Zone, owner entity.Entity, contents []entity.Entity) entity.Entity {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	z := g.EM.Create()
	z.AddComponent(uint(component.ZoneType))
	z.AddComponent(uint(component.Owner))
	z.AddComponent(uint(component.Contains))

	g.ZoneTypes[z.ID().Index] = component.ZoneComponent{Value: zoneType}
	g.Owners[z.ID().Index] = component.OwnerComponent{Value: owner.ID()}

	ids := make([]entity.EntityId, len(contents))
	for i, e := range contents {
		ids[i] = e.ID()
	}
	g.Contains[z.ID().Index] = component.ContainsComponent{Value: ids}

	return z
}

func (g *Game) SpawnPlayers(count int, initialLT int) {
	for i := 0; i < count; i++ {
		ent := g.SpawnWithLifeTotal(initialLT)
		g.Lock.Lock()
		g.Players = append(g.Players, ent)
		g.Lock.Unlock()
	}
}

func (g *Game) SpawnPlayer(initialLT int) entity.Entity {
	ent := g.SpawnWithLifeTotal(initialLT)
	g.Lock.Lock()
	g.Players = append(g.Players, ent)
	g.Lock.Unlock()
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
