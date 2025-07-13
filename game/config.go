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
	// required
	Name     string                  `toml:"name"`
	ZoneType component.ZoneComponent `toml:"zone_type,omitempty"`

	// optional components
	CardData *component.CardDataComponent `toml:"card_data,omitempty"`

	ManaCost       *component.ManaCostComponent       `toml:"mana_cost"`
	TypeLine       *component.TypeLineConfig          `toml:"type_line"`
	ColorIndicator *component.ColorIndicatorComponent `toml:"color_indicator"`
	Illustration   *component.IllustrationComponent   `toml:"illustration,omitempty"`

	// TODO
	// PowerAndToughness  *component.PowerAndToughnessComponent  `toml:"power_and_toughness,omitempty"`
	// Loyalty            *component.LoyaltyComponent            `toml:"loyalty,omitempty"`
	// Defense            *component.DefenseComponent            `toml:"defense,omitempty"`
	// HandModifier       *component.HandModifierComponent       `toml:"hand_modifier,omitempty"`
	// LifeModifier       *component.LifeModifierComponent       `toml:"life_modifier,omitempty"`
	// IllustrationCredit *component.IllustrationCreditComponent `toml:"illustration_credit,omitempty"`
	// LegalText          *component.LegalTextComponent          `toml:"legal_text,omitempty"`
	// CollectorNumber    *component.CollectorNumberComponent    `toml:"collector_number,omitempty"`
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
