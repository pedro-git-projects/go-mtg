package game

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/pedro-git-projects/go-mtg/component"
)

type LibraryConfig struct {
	Player int          `toml:"player"` // 1-based index
	Cards  []CardConfig `toml:"cards"`
}

type GameConfig struct {
	Players     int             `toml:"players"`
	InitialLife int             `toml:"initial_life"`
	Libraries   []LibraryConfig `toml:"libraries"`
}

type CardConfig struct {
	Name           string                            `toml:"name"`
	ManaCost       component.ManaCostComponent       `toml:"mana_cost"`
	ColorIndicator component.ColorIndicatorComponent `toml:"color_indicator"`
}

func LoadConfig(path string) (*GameConfig, error) {
	var cfg GameConfig
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if _, err := toml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}
	// sanity check: one [[libraries]] per player
	if cfg.Players != len(cfg.Libraries) {
		return nil, fmt.Errorf(
			"players = %d but %d [[libraries]] entries",
			cfg.Players, len(cfg.Libraries),
		)
	}
	return &cfg, nil
}
