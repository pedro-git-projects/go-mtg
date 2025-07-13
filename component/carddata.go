package component

import (
	"encoding/json"
	"fmt"
	"os"
)

type JSONCardData struct {
	Name              string  `json:"name"`
	ManaCost          string  `json:"mana_cost"`
	ConvertedManaCost float32 `json:"cmc"`
	TypeLine          string  `json:"type_line"`
	OracleText        string  `json:"oracle_text"`

	Power     string `json:"power,omitempty"`
	Toughness string `json:"toughness,omitempty"`
	Loyalty   string `json:"loyalty,omitempty"`

	Colors        []string `json:"colors"`
	ColorIdentity []string `json:"color_identity"`

	Keywords     []string `json:"keywords,omitempty"`
	ProducedMana []string `json:"produced_mana,omitempty"`
}

type CardDataComponent struct {
	Path string       `toml:"path"`
	Data JSONCardData `json:"-"`
}

func (c *CardDataComponent) UnmarshalTOML(v any) error {
	tbl, ok := v.(map[string]any)
	if !ok {
		return fmt.Errorf("invalid TOML for CardDataComponent: %T", v)
	}
	raw, ok := tbl["path"].(string)
	if !ok {
		return fmt.Errorf("carddata.path must be a string")
	}
	c.Path = raw
	return nil
}

func (c *CardDataComponent) LoadFromDisk(rootDir string) error {
	b, err := os.ReadFile(rootDir + "/" + c.Path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &c.Data)
}
