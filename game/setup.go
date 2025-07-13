package game

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pedro-git-projects/go-mtg/component"
	"github.com/pedro-git-projects/go-mtg/entity"
)

const jsonDir = "assets/json/cards_json"

func SetupStandardGameFromTOML(path string) (*Game, error) {
	// build card name -> .json lookup map
	if err := InitCardJSONLookup(jsonDir); err != nil {
		return nil, fmt.Errorf("initializing card lookup: %w", err)
	}

	cfg, err := LoadConfig(path)
	if err != nil {
		return nil, err
	}
	if cfg.Players != len(cfg.Libraries) {
		return nil, fmt.Errorf("players = %d but got %d library sections",
			cfg.Players, len(cfg.Libraries))
	}

	for li := range cfg.Libraries {
		for ci := range cfg.Libraries[li].Cards {
			card := &cfg.Libraries[li].Cards[ci]

			if card.CardData == nil {
				if fname, ok := LookupCardJSON(card.Name); ok {
					card.CardData = &component.CardDataComponent{Path: fname}
				} else {
					return nil, fmt.Errorf("no JSON found for card %q", card.Name)
				}
			}

			if err := card.CardData.LoadFromDisk(jsonDir); err != nil {
				return nil, fmt.Errorf("loading card JSON %q: %w",
					card.CardData.Path, err)
			}
		}
	}

	// estimate capacity & make the game
	totalCards := 0
	for _, lib := range cfg.Libraries {
		totalCards += len(lib.Cards)
	}
	g := NewGame(cfg.Players + totalCards + cfg.Players)

	// spawn each player and their library
	for pi, lib := range cfg.Libraries {
		player := g.SpawnWithLifeTotal(cfg.InitialLife)
		g.Players = append(g.Players, player)

		var deck []entity.Entity
		for _, rawCfg := range lib.Cards {
			patched := cardCfgToEntityCfg(rawCfg)
			c := g.SpawnCard(patched)
			deck = append(deck, c)
		}

		g.SpawnZone(component.Library, player, deck)
		fmt.Printf("Player %d: loaded %d cards\n", pi+1, len(deck))
	}

	return g, nil
}

// cardCfgToEntityCfg fills in any missing components from the JSON.
func cardCfgToEntityCfg(cfg CardConfig) CardConfig {
	data := cfg.CardData.Data

	// 1) Name
	if cfg.Name == "" {
		cfg.Name = data.Name
	}

	// 2) ManaCost
	if cfg.ManaCost == nil {
		if mc, err := component.ParseManaCost(data.ManaCost); err == nil {
			cfg.ManaCost = &mc
		}
	}

	// 3) TypeLine
	if cfg.TypeLine == nil {
		tl := component.ParseTypeLine(data.TypeLine)
		cfg.TypeLine = &tl
	}

	// 4) ColorIndicator
	if cfg.ColorIndicator == nil {
		raw := component.ColorFromStrings(data.Colors)
		ci := component.ColorIndicatorComponent{Value: raw}
		cfg.ColorIndicator = &ci
	}

	// 5) Illustration key <- take JSON filename without “.json”
	if cfg.Illustration == nil {
		base := filepath.Base(cfg.CardData.Path)            // "Abattoir_Ghoul_85.json"
		key := strings.TrimSuffix(base, filepath.Ext(base)) // "Abattoir_Ghoul_85"
		cfg.Illustration = &component.IllustrationComponent{AssetKey: key}
	}

	return cfg
}
